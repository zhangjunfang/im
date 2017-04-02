package tfClient

import (
	"errors"
	"fmt"
	"runtime/debug"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/donnie4w/go-logger/logger"
	"github.com/zhangjunfang/im/protocol"
)

func HttpClient(f func(*protocol.ImClient) error, urlstr string) (err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprint(er))
			logger.Error(er)
			logger.Error(string(debug.Stack()))
		}
	}()
	if urlstr != "" {
		transport, err := thrift.NewTHttpPostClient(urlstr)
		defer transport.Close()
		if err == nil {
			factory := thrift.NewTCompactProtocolFactory()
			transport.Open()
			imClient := protocol.NewITimClientFactory(transport, factory)
			err = f(imClient)
		}
	}
	return
}

func HttpClient2(f func(*protocol.ImClient) error, user_auth_url string) (err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprint(er))
			logger.Error(er)
			logger.Error(string(debug.Stack()))
		}
	}()
	if user_auth_url != "" {
		transport, err := thrift.NewTHttpPostClient(user_auth_url)
		defer transport.Close()
		if err == nil {
			factory := thrift.NewTCompactProtocolFactory()
			transport.Open()
			itimClient := protocol.NewITimClientFactory(transport, factory)
			err = f(itimClient)
		}
	} else {
		err = errors.New("httpclient url is null")
	}
	return
}

func TcpClient(f func(*protocol.ImClient), urlstr string) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(err)
		}
	}()
	if urlstr != "" {
		transport, err := thrift.NewTSocket(urlstr)
		defer transport.Close()
		if err == nil {
			protocolFactory := thrift.NewTCompactProtocolFactory()
			itimClient := protocol.NewITimClientFactory(transport, protocolFactory)
			f(itimClient)
		}
	}
}
