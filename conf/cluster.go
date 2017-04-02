package conf

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"runtime/debug"
	"strconv"

	"github.com/donnie4w/dom4g"
	"github.com/donnie4w/go-logger/logger"
)

/**配置结构对象*/

type ClusterBean struct {
	RedisAddr   string //  redis ip:port
	RedisPwd    string //
	RedisDB     int    //
	RequestAddr string // 访问地址
	RequestType string // 访问类型
	Domain      string // 域名
	Username    string // 登陆名
	Password    string // 登陆密码
	IsCluster   int    // 1集群 0不集群
	Interflow   int    // 合流信息发送 0不合流  1合流
	Keytimeout  int    // key 过期时间
}

func (cb *ClusterBean) Init(filexml string) (b bool) {
	//异常处理
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Init error", err)
			logger.Error(string(debug.Stack()))
			b = false
		}
	}()
	//判断文件不存在
	if !isExist(filexml) {
		return false
	}
	//打开文件
	xmlconfig, err := os.Open(filexml)
	//异常处理
	if err != nil {
		panic(fmt.Sprint("xmlconfig is error:", err.Error()))
		os.Exit(0)
	}
	config, err := ioutil.ReadAll(xmlconfig)
	if err != nil {
		panic(fmt.Sprint("config is error:", err.Error()))
		os.Exit(1)
	}
	//使用dom4g解析xml
	dom, err := dom4g.LoadByXml(string(config))
	if err == nil {
		nodes := dom.AllNodes()
		if nodes != nil {
			for _, node := range nodes {
				name := node.Name()
				value := node.Value
				v := reflect.ValueOf(cb).Elem().FieldByName(name)
				if v.CanSet() {
					switch v.Type().Name() {
					case "string":
						v.Set(reflect.ValueOf(value))
					case "int":
						i, _ := strconv.Atoi(value)
						v.Set(reflect.ValueOf(i))
					default:
						fmt.Println("other type:", v.Type().Name(), ">>>", name)
					}
				}
			}
		}
	}
	b = true
	return
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
