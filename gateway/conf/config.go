package conf

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	log "github.com/shenjing023/llog"
	"gopkg.in/yaml.v3"
)

// YamlConf global config struct
type YamlConf struct {
	Debug   bool   `yaml:"debug"`
	Port    int    `yaml:"port"`
	LogDir  string `yaml:"log_dir"`
	Account struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"account"`
	Post struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"post"`
	AuthToken struct {
		Access struct {
			SignKey string `yaml:"sign_key"`
			Expire  int    `yaml:"expire"`
		} `yaml:"access"`
		Refresh struct {
			SignKey string `yaml:"sign_key"`
			Expire  int    `yaml:"expire"`
		} `yaml:"refresh"`
	} `yaml:"auth_token"`
	Salt      string `yaml:"salt"`
	QueryDeep int    `yaml:"query_deep"`
}

// Cfg global config variate
var Cfg = new(YamlConf)

// Init init global config
func Init(configPath string) {
	//加载配置文件
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println("Open config file error:", err.Error())
		os.Exit(1)
	}
	if err = yaml.Unmarshal(file, Cfg); err != nil {
		fmt.Println("Read yaml config file error:", err.Error())
		os.Exit(1)
	}
	if Cfg.Debug {
		if len(Cfg.LogDir) > 0 {
			// set file log
			log.SetFileLogger(
				path.Join(Cfg.LogDir, "service.log")+"-%Y%m%d%H%M",
				log.WithCaller(true),
				log.WithMaxAge(7*24*time.Hour),
				log.WithRotationTime(24*time.Hour),
				log.WithLevel(log.DebugLevel),
			)
		} else {
			log.SetConsoleLogger(
				log.WithCaller(true),
				log.WithLevel(log.DebugLevel),
			)
		}
	} else {
		log.SetConsoleLogger(
			log.WithColor(false),
			log.WithJSON(true),
		)
	}
}
