package dao

import (
	"changweiba-backend/conf"
	"fmt"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
)

var dbEngine *xorm.Engine

func Init(){
	var err error
	dataSourceName:=fmt.Sprintf("%s:%s@%s:%s/%s",conf.Cfg.DB.User,conf.Cfg.DB.Passwd,conf.Cfg.DB.Host, 
		conf.Cfg.DB.Port,conf.Cfg.DB.Dbname)
	dbEngine, err = xorm.NewEngine("mysql", dataSourceName)
	if err!=nil{
		os.Exit(1)
		logs.Error(err.Error())
	}
}
