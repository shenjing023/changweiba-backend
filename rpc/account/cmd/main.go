package main

import (
	"changweiba-backend/conf"
	"changweiba-backend/rpc/account"
	"flag"
	"fmt"
	"os"
	"runtime"
)

func main() {
	//设置可同时使用的CPU数目，
	runtime.GOMAXPROCS(runtime.NumCPU())
	//命令行解析
	pwd, _ := os.Getwd()
	execDir := flag.String("d", pwd, "execute directory")
	flag.Parse()
	fmt.Println("Current execute directory is:", *execDir)
	conf.InitConfig(*execDir)
	account.NewAccountService("localhost", conf.Cfg.Account.Port)
}
