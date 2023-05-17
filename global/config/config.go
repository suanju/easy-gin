package config

import (
	"log"
	"path/filepath"

	"gopkg.in/ini.v1"
)

type Info struct {
	SqlConfig *SqlConfigStruct
	RConfig   *RConfigStruct
}

func init() {
	//避免全局重复导包
	ReturnsInstance()
}

var Config = new(Info)
var cfg *ini.File
var err error

type SqlConfigStruct struct {
	IP       string `ini:"ip"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Host     int    `ini:"host"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type RConfigStruct struct {
	IP       string `ini:"ip"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
}

func ReturnsInstance() *Info {
	Config.SqlConfig = &SqlConfigStruct{}
	cfg, err = ini.Load(filepath.ToSlash("./conf/app.ini"))
	if err != nil {
		log.Fatalf("配置文件不存在,请检查环境: %v \n", err)
	}

	err = cfg.Section("mysql").MapTo(Config.SqlConfig)
	if err != nil {
		log.Fatalf("Mysql读取配置文件错误: %v \n", err)
	}
	Config.RConfig = &RConfigStruct{}
	err = cfg.Section("redis").MapTo(Config.RConfig)
	if err != nil {
		log.Fatalf("Redis读取配置文件错误: %v \n", err)
	}
	return Config
}
