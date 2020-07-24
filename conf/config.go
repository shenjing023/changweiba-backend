package conf

import (
	"changweiba-backend/pkg/logs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type YamlConf struct {
	Debug   bool `yaml:"debug"`
	Graphql struct {
		Port       int    `yaml:"port"`
		JwtSignkey string `yaml:"jwt_signkey"`
	} `yaml:"graphql,omitempty"`
	Account struct {
		Port int `yaml:"port"`
	} `yaml:"account,omitempty"`
	Post struct {
		Port int `yaml:"port"`
	} `yaml:"post,omitempty"`
	DB struct {
		Host      string `yaml:"host"`
		Port      int    `yaml:"port"`
		User      string `yaml:"user"`
		Password  string `yaml:"password"`
		Dbname    string `yaml:"dbname"`
		MaxIdle   int    `yaml:"max_idle,omitempty"`
		MaxOpen   int    `yaml:"max_open,omitempty"`
		LogConfig string
	} `yaml:"db"`
	SignKey   string `yaml:"sign_key"`
	Salt      string `yaml:"salt"`
	QueryDeep int    `yaml:"query_deep"`

	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
}

//beego日志配置结构体
type LoggerConfig struct {
	FileName            string `json:"filename"`
	Level               int    `json:"level"`    //日志保存的时候的级别，默认是 Trace 级别
	Maxlines            int    `json:"maxlines"` //每个文件保存的最大行数，默认值 1000000
	Maxsize             int    `json:"maxsize"`  //每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
	Daily               bool   `json:"daily"`    //是否按照每天logrotate，默认是 true
	Maxdays             int    `json:"maxdays"`  //文件最多保存多少天，默认保存 7 天
	Rotate              bool   `json:"rotate"`   //是否开启 logrotate，默认是 true
	Perm                string `json:"perm"`     //日志文件权限
	RotatePerm          string `json:"rotateperm"`
	EnableFuncCallDepth bool   `json:"-"` //输出文件名和行号
	LogFuncCallDepth    int    `json:"-"` //函数调用层级
}

var Cfg *YamlConf

func init() {
	Cfg = new(YamlConf)
}

func InitConfig(executeDir string) {
	//加载配置文件
	file, err := ioutil.ReadFile(executeDir + "/config.yaml")
	if err != nil {
		fmt.Println("Read config file error:", err.Error())
		os.Exit(1)
	}
	if err = yaml.Unmarshal(file, Cfg); err != nil {
		fmt.Println("Read yaml config file error:", err.Error())
		os.Exit(1)
	}

	//db日志配置
	var dbLogConf = LoggerConfig{
		FileName:            executeDir + "/log/sql.log",
		Level:               7,
		EnableFuncCallDepth: true,
		LogFuncCallDepth:    3,
		RotatePerm:          "777",
		Perm:                "777",
		Daily:               true,
		Rotate:              true,
		Maxdays:             30,
	}
	cfg, _ := json.Marshal(&dbLogConf)
	Cfg.DB.LogConfig = string(cfg)

	//服务日志配置
	var logConf = LoggerConfig{
		FileName:            executeDir + "/log/log.log",
		Level:               7,
		EnableFuncCallDepth: true,
		LogFuncCallDepth:    3,
		RotatePerm:          "777",
		Perm:                "777",
		Daily:               true,
		Rotate:              true,
		Maxdays:             30,
	}
	cfg, _ = json.Marshal(&logConf)
	if Cfg.Debug {
		_ = logs.SetLogger("console")
	}
	if err = logs.SetLogger(logs.AdapterFile, string(cfg)); err != nil {
		fmt.Println("Init server logger error:", err.Error())
		os.Exit(1)
	}
	logs.EnableFuncCallDepth(logConf.EnableFuncCallDepth)
	logs.SetLogFuncCallDepth(logConf.LogFuncCallDepth)
}
