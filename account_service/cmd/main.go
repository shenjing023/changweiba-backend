package main

import (
	service "cw_account_service"
	"flag"
	"os"
)

func main() {
	pwd, _ := os.Getwd()
	execDir := flag.String("d", pwd, "execute directory")
	flag.Parse()
	service.RunAccountService(*execDir + "/conf/config.yaml")
	//conf.InitConfig()
}
