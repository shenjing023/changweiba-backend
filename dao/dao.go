package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"changweiba-backend/conf"
	"os"
	"github.com/astaxie/beego/logs"
)

var dbEngine *xorm.Engine

func init(){
	var err error
	dbEngine, err = xorm.NewEngine("mysql", "root:123@/test?charset=utf8")
	if err!=nil{
		os.Exit(1)
		logs.Error(err.Error())
	}
}
