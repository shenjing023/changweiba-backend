package conf

import (
	"fmt"
	"log/slog"
	"os"

	llog "github.com/shenjing023/vivy-polaris/log"
	"gopkg.in/yaml.v3"
)

// YamlConf global config struct
type YamlConf struct {
	Debug bool `yaml:"debug"`
	Port  int  `yaml:"port"`
	DB    struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Dbname   string `yaml:"dbname"`
		MaxIdle  int    `yaml:"max_idle,omitempty"` //设置连接池中空闲连接的最大数量
		MaxOpen  int    `yaml:"max_open,omitempty"` //设置打开数据库连接的最大数量
	} `yaml:"db"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
	Etcd struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"etcd"`
	LogDir           string `yaml:"log_dir"`
	Salt             string `yaml:"salt"`
	JaegerCollectURL string `yaml:"jeager_collect_url"`
}

// Cfg global config variate
var Cfg = new(YamlConf)

// Init init global config
func Init(configPath string) {
	//加载配置文件
	file, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Open config file error:", err.Error())
		os.Exit(1)
	}
	if err = yaml.Unmarshal(file, Cfg); err != nil {
		fmt.Println("Read yaml config file error:", err.Error())
		os.Exit(1)
	}
	if Cfg.Debug {
		llog.Init(llog.WithLevel(slog.LevelDebug))
	} else {
		llog.Init()
	}
}
