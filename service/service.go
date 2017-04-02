package service

import (
	"fmt"
	"net"

	"github.com/donnie4w/go-logger/logger"
	"github.com/zhangjunfang/im/common"
)

func Start() {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	service := fmt.Sprint(common.CF.Addr, ":", common.CF.Port)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err == nil {
			go handler(conn)
		}
	}
}

func handler(conn net.Conn) {
	fmt.Println("handler")
}

func checkError(err error) {
	if err != nil {
		logger.Error(err.Error())
		panic(err.Error())
	}
}
