package service

import (
	"crypto/tls"
	"errors"
	"fmt"
	"runtime/debug"
	"strconv"
	"sync"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/donnie4w/go-logger/logger"
	"github.com/zhangjunfang/im/cluster"
	"github.com/zhangjunfang/im/clusterRoute"
	"github.com/zhangjunfang/im/clusterServer"
	"github.com/zhangjunfang/im/common"
	"github.com/zhangjunfang/im/connect"
	"github.com/zhangjunfang/im/fw"
	"github.com/zhangjunfang/im/impl"
	"github.com/zhangjunfang/im/protocol"
	"github.com/zhangjunfang/im/route"
	"github.com/zhangjunfang/im/thriftserver"
	"github.com/zhangjunfang/im/utils"
)

type Controlloer struct {
	Port int
	Ip   string
}

func ServerStart() {
	go Httpserver()
	go tsslServer()
	s := new(Controlloer)
	s.SetAddr(common.CF.GetIp())
	if cluster.IsCluster() {
		go clusterServer.ServerStart()
	}
	s.Server()
}

func (t *Controlloer) SetAddr(port int, ip string) {
	t.Port = port
	t.Ip = ip
}

func (t *Controlloer) ListenAddr() string {
	return fmt.Sprint(t.Ip, ":", t.Port)
}

func (t *Controlloer) Server() {
	transportFactory := thrift.NewTBufferedTransportFactory(1024)
	protocolFactory := thrift.NewTCompactProtocolFactory()
	serverTransport, err := thrift.NewTServerSocket(t.ListenAddr())
	if err != nil {
		logger.Error("server:", err.Error())
		panic(err.Error())
	}
	handler := new(impl.TimImpl)
	processor := protocol.NewITimProcessor(handler)
	server := thriftserver.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("server listen:", t.ListenAddr())
	Listen(server, 100)
	if err == nil {
		for {
			client, err := Accept(server)
			if err == nil {
				go controllerHandler(client)
			}
		}
	}
}

func Listen(server *thriftserver.TSimpleServer, count int) (err error) {
	if count <= 0 {
		err = errors.New("")
		return
	}
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Listen,", err)
			logger.Error(string(debug.Stack()))
			count--
			Listen(server, count)
		}
	}()
	err = server.Listen()
	return
}

func Accept(server *thriftserver.TSimpleServer) (client thrift.TTransport, err error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Accept,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	client, err = server.Accept()
	return
}

func tsslServer() {
	if common.CF.TLSPort <= 0 && common.CF.TLSServerPem == "" && common.CF.TLSServerKey == "" {
		return
	}
	cer, err := tls.LoadX509KeyPair(common.CF.TLSServerPem, common.CF.TLSServerKey)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	tsslServerSocket, err := thrift.NewTSSLServerSocket(fmt.Sprint(common.CF.Addr, ":", common.CF.TLSPort), config)
	if err != nil {
		return
	}
	err = tsslServerSocket.Listen()
	if err != nil {
		return
	}
	fmt.Println("tls server listen:", common.CF.TLSPort)
	for {
		client, err := tsslServerSocket.Accept()
		if err == nil && client != nil {
			go controllerHandler(client)
		}
	}
}

func controllerHandler(tt thrift.TTransport) {
	isclose := false
	var gorutineclose *bool = &isclose
	defer func() {
		if err := recover(); err != nil {
			logger.Error("controllerHandler,", err)
			*gorutineclose = true
		}
	}()
	tu := &connect.TimUser{Client: NewTimClient(tt), OverLimit: 3, Fw: fw.CONNECT, IdCardNo: utils.TimeMills(), Sendflag: make(chan string, 0), Sync: new(sync.Mutex)}
	connect.TP.AddConnect(tu)
	defer func() {
		if cluster.IsCluster() && tu.UserTid != nil {
			loginname, err := connect.GetLoginName(tu.UserTid)
			if loginname != "" && err == nil {
				cluster.DelLoginnameFromCluter(loginname)
			}
			if common.CF.Presence == 1 {
				p := impl.OfflinePBean(tu.UserTid)
				go clusterRoute.ClusterRoutePBean(p)
			} else {
				go route.RoutePBean(impl.OfflinePBean(tu.UserTid))
			}
		} else if tu.UserTid != nil && common.CF.Presence == 1 {
			go route.RoutePBean(impl.OfflinePBean(tu.UserTid))
		}
		connect.TP.DeleteTimUser(tu)
	}()
	defer func() { tt.Close() }()
	monitorChan := make(chan string, 2)
	heartbeat := common.CF.HeartBeat
	if heartbeat == 0 {
		heartbeat = 30 * 60
	}
	//	if heartbeat > 0 {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				//				logger.Error(string(debug.Stack()))
				logger.Error(err)
			}
		}()
		defer func() {
			if err := recover(); err != nil {
			}
			*gorutineclose = true
			monitorChan <- "ping end"
		}()
		checkinCluster := 0
		for {
			if *gorutineclose {
				break
			}
			for i := 0; i < heartbeat; i++ {
				time.Sleep(1 * time.Second)
				if *gorutineclose {
					goto END
				}
				if tu.OverLimit <= 0 {
					goto END
				}
				if tu.Fw == fw.CLOSE {
					goto END
				}
				checkinCluster++
				if checkinCluster >= common.ClusterConf.Keytimeout/3 {
					checkinCluster = 0
					if cluster.IsCluster() && tu.UserTid != nil {
						loginname, err := connect.GetLoginName(tu.UserTid)
						if loginname != "" && err == nil {
							cluster.SetLoginnameToCluster(loginname)
						}
					}
				}
			}
			if tu.Fw == fw.AUTH {
				er := tu.Ping()
				if er != nil {
					logger.Error("ping err", er.Error())
					panic("ping err")
				}
				tu.OverLimit--
			} else {

				goto END
			}
			if tu.Fw == fw.CLOSE {
				break
			}
		}
	END:
	}()
	//	}
	go TimProcessor(tt, tu, gorutineclose, monitorChan)
	<-monitorChan
}

func NewTimClient(tt thrift.TTransport) *protocol.ImClient {
	transportFactory := thrift.NewTBufferedTransportFactory(1024)
	protocolFactory := thrift.NewTCompactProtocolFactory()
	useTransport := transportFactory.GetTransport(tt)
	return protocol.NewITimClientFactory(useTransport, protocolFactory)
}

func TimProcessor(client thrift.TTransport, tu *connect.TimUser, gorutineclose *bool, monitorChan chan string) error {
	defer func() {
		if err := recover(); err != nil {
			//			logger.Error(string(debug.Stack()))
			logger.Warn("processor:", err)
		}
	}()
	defer func() {
		if err := recover(); err != nil {
			//			logger.Error("TimProcessor error", err)
			//			logger.Error(string(debug.Stack()))
			logger.Warn("processor:", err)
		}
		*gorutineclose = true
		monitorChan <- "timProcessor end"
	}()
	compactprotocol := thrift.NewTCompactProtocol(client)
	pub := strconv.Itoa(time.Now().Nanosecond())
	handler := &impl.TimImpl{Pub: pub, Client: client, Tu: tu}
	processor := protocol.NewITimProcessor(handler)
	for {
		ok, err := processor.Process(compactprotocol, compactprotocol)
		if err, ok := err.(thrift.TTransportException); ok && err.TypeId() == thrift.END_OF_FILE {
			return nil
		} else if err != nil {
			return err
		}
		if !ok {
			logger.Error("Processor error:", err)
			break
		}
		if tu.Fw == fw.CLOSE || tu.OverLimit <= 0 {
			break
		}
	}
	return nil
}
