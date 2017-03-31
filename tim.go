package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zhangjunfang/im/cluster"
	"github.com/zhangjunfang/im/myDb"

	gdao "github.com/donnie4w/gdao"
	"github.com/donnie4w/go-logger/logger"
	. "github.com/zhangjunfang/im/common"
	"github.com/zhangjunfang/im/daoService"
	. "github.com/zhangjunfang/im/protocol"
	"github.com/zhangjunfang/im/service"
	"github.com/zhangjunfang/im/ticker"
)

func init() {
	servername := fmt.Sprint("im", ProtocolversionName, " server")
	fmt.Println("----------------------------------------------------------")
	fmt.Println("-------------------- " + servername + " ---------------------")
	fmt.Println("--------------------------------------------------------")
	fmt.Println("------------------ zhangjunfang0505@163.com ------------------")
	fmt.Println("----------------------------------------------------------")
}

func initGdao() {
	if CF.Db_Exsit == 0 {
		return
	}
	logger.Debug("initGdao")
	myDb.Init()
	gdao.SetDB(myDb.Master)
	gdao.SetAdapterType(gdao.MYSQL)
	gbs, err := gdao.ExecuteQuery("select 1")
	if err == nil {
		logger.Debug("test db ok", gbs[0].MapIndex(1).Value())
	}
}

func initLog(loglevel string) {
	logger.SetConsole(true)
	logger.SetRollingDaily(CF.GetLog())
	switch loglevel {
	case "debug":
		logger.SetLevel(logger.DEBUG)
	case "info":
		logger.SetLevel(logger.INFO)
	case "warn":
		logger.SetLevel(logger.WARN)
	case "error":
		logger.SetLevel(logger.ERROR)
	default:
		logger.SetLevel(logger.WARN)
	}
}

//tim f tim.xml c cluster.xml d debug
func main() {
	flag.Parse()
	wd, _ := os.Getwd()
	if flag.NArg() > 6 {
		fmt.Println("error:", "flag's length is", flag.NArg())
		os.Exit(1)
	}
	timconf := fmt.Sprint(wd, "/tim.xml")
	initconf := ""
	clusterconf := fmt.Sprint(fmt.Sprint(wd, "/cluster.xml"))
	for i := 0; i < flag.NArg(); i++ {
		if i%2 == 0 {
			switch flag.Arg(i) {
			case "f":
				timconf = flag.Arg(i + 1)
			case "c":
				clusterconf = flag.Arg(i + 1)
			case "d":
				initconf = flag.Arg(i + 1)
			default:
				fmt.Println("error:", "error arg:", flag.Arg(i))
				os.Exit(1)
			}
		}
	}
	CF.Init(timconf)
	initLog(initconf)
	cluster.InitCluster(clusterconf)
	initGdao()
	daoService.InitDaoservice()
	ticker.TickerStart()
	service.ServerStart()
}
