package dao

import (
	"changweiba-backend/conf"
	"fmt"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"math/big"
	"net"
	"os"
	"time"
)

var dbEngine *xorm.Engine

func Init(){
	var err error
	dataSourceName:=fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",conf.Cfg.DB.User,conf.Cfg.DB.Passwd,
		conf.Cfg.DB.Host, 
		conf.Cfg.DB.Port,conf.Cfg.DB.Dbname)
	dbEngine, err = xorm.NewEngine("mysql", dataSourceName)
	if err!=nil{
		logs.Error(err.Error())
		os.Exit(1)
	}
	if err=dbEngine.Ping();err!=nil{
		logs.Error(err.Error())
		os.Exit(1)
	}
	if conf.Cfg.DB.MaxIdleConns>0{
		dbEngine.SetMaxIdleConns(conf.Cfg.DB.MaxIdleConns)
	}
	if conf.Cfg.DB.MaxOpenConns>0{
		dbEngine.SetMaxOpenConns(conf.Cfg.DB.MaxOpenConns)
	}
	//日志
	f,err:=os.Create(conf.Cfg.DB.LogFile)
	if err!=nil{
		logs.Error(err.Error())
	}
	dbEngine.SetLogger(xorm.NewSimpleLogger(f))
	dbEngine.ShowSQL(true)	//不能忽略
}

func InsertUser(userName,password,ip string) (int64,error){
	now:=time.Now().Unix()
	user:=User{
		Name:userName,
		Password:password,
		Ip:InetAtoi(ip),
		CreateTime:now,
		LastUpdate:now,
	}
	_,err:=dbEngine.InsertOne(&user)
	if err!=nil{
		return 0,err
	}
	//spew.Dump(user)
	return user.Id,nil
}

//ip地址int->string相互转换
func InetAtoi(ip string) int64{
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

func InetItoa(ip int64) string{
	return fmt.Sprintf("%d.%d.%d.%d", byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}
