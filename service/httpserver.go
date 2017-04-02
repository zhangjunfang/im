package service

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime/debug"
	"strings"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"

	"github.com/donnie4w/go-logger/logger"
	"github.com/zhangjunfang/im/common"
	"github.com/zhangjunfang/im/daoService"
	"github.com/zhangjunfang/im/impl"
	"github.com/zhangjunfang/im/protocol"
)

func Httpserver() {
	if common.CF.GetHttpPort() <= 0 {
		return
	}
	http.HandleFunc("/tim", tim)
	http.HandleFunc("/info", info)
	http.HandleFunc("/uinfo", userInfo)
	http.HandleFunc("/hi", hbaseclient)
	s := &http.Server{
		Addr:           fmt.Sprint(":", common.CF.GetHttpPort()),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func tim(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("err:", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	X_Forwarded_For := r.Header.Get("X-Forwarded-For")
	ss := strings.Split(r.RemoteAddr, ":")
	ipaddr := ss[0]
	if X_Forwarded_For != "" && X_Forwarded_For != "127.0.0.1" {
		ipaddr = X_Forwarded_For
	}
	if r.ContentLength >= 100*1024*1024 {
		return
	}
	if !daoService.AllowHttpIp(ipaddr) {
		return
	}
	if "POST" == r.Method {
		protocolFactory := thrift.NewTCompactProtocolFactory()
		transport := thrift.NewStreamTransport(r.Body, w)
		inProtocol := protocolFactory.GetProtocol(transport)
		outProtocol := protocolFactory.GetProtocol(transport)
		handler := &impl.TimImpl{Ip: ipaddr}
		processor := protocol.NewITimProcessor(handler)
		processor.Process(inProtocol, outProtocol)
	}
}
