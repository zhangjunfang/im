package myDb

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zhangjunfang/im/common"
)

var Master *sql.DB

func Init() {
	initmaster()
}

//数据库连接管理
func initmaster() {
	if Master == nil {
		//获取数据库初始化参数
		dataSourceName, maxOpenConns, maxIdleConns := common.CF.GetDB()
		var err error
		//获取数据库连接
		Master, err = GetDB(dataSourceName, maxOpenConns, maxIdleConns)
		if err != nil {
			fmt.Println("any error on open database ", err.Error())
			os.Exit(1)
		}
	}
}

func GetDB(dataSourceName string, maxOpenConns, maxIdleConns int) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", dataSourceName)
	if err == nil {
		db.SetMaxOpenConns(maxOpenConns)
		db.SetMaxIdleConns(maxIdleConns)
	}
	return db, err
}
