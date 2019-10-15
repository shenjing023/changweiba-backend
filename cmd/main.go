package main

import (
	"changweiba-backend/conf"
	"changweiba-backend/dao"
	"changweiba-backend/rpc/account"
	"changweiba-backend/rpc/post"
	"flag"
	"fmt"
	"github.com/astaxie/beego/logs"
	"os"
	"runtime"
)

func main(){
	//设置可同时使用的CPU数目，
	runtime.GOMAXPROCS(runtime.NumCPU())
	//命令行解析
	pwd, _ := os.Getwd()
	execDir := flag.String("d", pwd, "execute directory")
	flag.Parse()
	fmt.Println("Current execute directory is:", *execDir)
	conf.InitConfig(*execDir)
	//初始化数据库
	dao.Init()
	//user服务
	logs.Info("account service port:", conf.Cfg.Account.Port)
	go account.NewAccountService("localhost",conf.Cfg.Account.Port)
	logs.Info("post service port:", conf.Cfg.Post.Port)
	go post.NewPostService("localhost",conf.Cfg.Post.Port)
	select {
	
	}
}
