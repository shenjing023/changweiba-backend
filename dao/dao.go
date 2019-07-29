package dao

import (
	"changweiba-backend/conf"
	"fmt"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
)

var DbEngine *xorm.Engine

func Init(){
	var err error
	dataSourceName:=fmt.Sprintf("%s:%s@%s:%s/%s",conf.Cfg.DB.User,conf.Cfg.DB.Passwd,conf.Cfg.DB.Host, 
		conf.Cfg.DB.Port,conf.Cfg.DB.Dbname)
	DbEngine, err = xorm.NewEngine("mysql", dataSourceName)
	if err!=nil{
		logs.Error(err.Error())
		os.Exit(1)
	}
	if conf.Cfg.DB.MaxIdleConns>0{
		DbEngine.SetMaxIdleConns(conf.Cfg.DB.MaxIdleConns)
	}
	if conf.Cfg.DB.MaxOpenConns>0{
		DbEngine.SetMaxOpenConns(conf.Cfg.DB.MaxOpenConns)
	}
	//日志
	if conf.Cfg.Debug{
		DbEngine.ShowSQL(true)
	}
	f,_:=os.Create(conf.Cfg.DB.LogFile)
	DbEngine.SetLogger(xorm.NewSimpleLogger(f))
}
