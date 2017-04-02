package common

import "github.com/zhangjunfang/im/conf"

/*版本*/
var VersionName = "im 1.0"
var VersionCode = 2
var Author = "ocean"
var Email = "zhangjunfang0505@163.com"

var CF = &conf.ConfBean{KV: make(map[string]string, 0), Db_Exsit: 1, MustAuth: 1}

var ClusterConf = &conf.ClusterBean{IsCluster: 1}
