package model

import (
	"github.com/donnie4w/gdao"
	"github.com/zhangjunfang/im/DB"
)

func init() {
	gdao.SetDB(DB.Master)
	gdao.SetAdapterType(gdao.MYSQL)
}
