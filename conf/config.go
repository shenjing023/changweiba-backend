package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type YamlConf struct {
	Graphql struct {
		Port       int    `yaml:"port"`
		JwtSignkey string `yaml:"jwt_signkey"`
	} `yaml:"graphql,omitempty"`
	Account struct {
		Port int `yaml:"port"`
	} `yaml:"account,omitempty"`
}

var Cfg *YamlConf

func InitConfig(executeDir string) {
	//加载配置文件
	file, err := ioutil.ReadFile(executeDir + "/config.yaml")
	if err != nil {
		fmt.Println("Open config file error:", err.Error())
		os.Exit(1)
	}
	Cfg = new(YamlConf)
	if err = yaml.Unmarshal(file, Cfg); err != nil {
		fmt.Println("Read yaml config file error:", err.Error())
		os.Exit(1)
	}
}
