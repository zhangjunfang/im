package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zhangjunfang/im/cluster"
	"github.com/zhangjunfang/im/common"
	"github.com/zhangjunfang/im/myDb"

	"github.com/donnie4w/gdao"
	"github.com/donnie4w/go-logger/logger"
	"github.com/zhangjunfang/im/daoService"
	"github.com/zhangjunfang/im/protocol"
	"github.com/zhangjunfang/im/service"
	"github.com/zhangjunfang/im/ticker"
)

//服务器版本相关的信息
func init() {
	servername := fmt.Sprint("im", protocol.ProtocolversionName, " server")
	fmt.Println("----------------------------------------------------------")
	fmt.Println("-------------------- " + servername + " ---------------------")
	fmt.Println("--------------------------------------------------------")
	fmt.Println("------------------ zhangjunfang0505@163.com ------------------")
	fmt.Println("----------------------------------------------------------")
}

//初始化数据访问层
func initGdao() {
	if common.CF.Db_Exsit == 0 {
		return
	}
	myDb.Init()                               //获取数据库初始化参数
	gdao.SetDB(myDb.Master)                   //设置数据库连接
	gdao.SetAdapterType(gdao.MYSQL)           //数据库类型设置
	gbs, err := gdao.ExecuteQuery("select 1") //测试数据库连接是否
	if err == nil {
		logger.Debug("test db ok", gbs[0].MapIndex(1).Value())
	}
}

//日志初始化  同时在控制台和 配置文件中位置 生产日志文件
func initLog(loglevel string) {
	logger.SetConsole(true)
	logger.SetRollingDaily(common.CF.GetLog())
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

//im f im.xml c cluster.xml d debug
func main() {
	//解析命令行参数
	flag.Parse()
	//获得当前工作路径
	wd, _ := os.Getwd()
	//判断当前命令行参数长度
	if flag.NArg() > 6 {
		fmt.Println("error:", "flag's length is", flag.NArg())
		os.Exit(1)
	}
	//默认im.xml配置文件相对路径的绝对位置
	imconf := fmt.Sprint(wd, "/im.xml")
	//日志级别  命令行初始化配置参数
	initconf := ""
	//默认cluster.xml集群配置文件位置
	clusterconf := fmt.Sprint(fmt.Sprint(wd, "/cluster.xml"))
	for i := 0; i < flag.NArg(); i++ {
		if i%2 == 0 {
			switch flag.Arg(i) {
			case "f":
				imconf = flag.Arg(i + 1)
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
	//初始化im配置文件，并解析
	common.CF.Init(imconf)
	//日志初始化  并使用控制和文件两种通道生成文件
	initLog(initconf)
	//集群配置解析
	cluster.InitCluster(clusterconf)
	//初始化数据访问层
	initGdao()
	//初始化数据服务层=============================
	daoService.InitDaoservice()
	//
	ticker.TickerStart()
	//
	service.ServerStart()
}
