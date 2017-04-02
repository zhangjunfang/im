package service

import (
	"fmt"
	"io"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/donnie4w/go-logger/logger"
	"github.com/zhangjunfang/im/connect"
	"github.com/zhangjunfang/im/hbase"
)

func info(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(debug.Stack())
		}
	}()
	X_Forwarded_For := r.Header.Get("X-Forwarded-For")
	ss := strings.Split(r.RemoteAddr, ":")
	ipaddr := ss[0]
	if X_Forwarded_For != "" && X_Forwarded_For != "127.0.0.1" {
		ipaddr = X_Forwarded_For
	}
	logger.Debug("ip:", ipaddr)
	if r.ContentLength >= 2*1024*1024 {
		return
	}
	str := fmt.Sprintln("user:", connect.TP.Len4P(), "===>", connect.TP.Len4PU())
	io.WriteString(w, str)
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(debug.Stack())
		}
	}()
	X_Forwarded_For := r.Header.Get("X-Forwarded-For")
	ss := strings.Split(r.RemoteAddr, ":")
	ipaddr := ss[0]
	if X_Forwarded_For != "" && X_Forwarded_For != "127.0.0.1" {
		ipaddr = X_Forwarded_For
	}
	logger.Debug("ip:", ipaddr)
	if r.ContentLength >= 2*1024*1024 {
		return
	}
	io.WriteString(w, connect.TP.PrintUsersInfo())
}

func hbaseclient(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(debug.Stack())
		}
	}()
	X_Forwarded_For := r.Header.Get("X-Forwarded-For")
	ss := strings.Split(r.RemoteAddr, ":")
	ipaddr := ss[0]
	if X_Forwarded_For != "" && X_Forwarded_For != "127.0.0.1" {
		ipaddr = X_Forwarded_For
	}
	logger.Debug("ip:", ipaddr)
	if r.ContentLength >= 2*1024*1024 {
		return
	}
	io.WriteString(w, hbase.PrintPoolInfo())
}
