package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Info struct {
	SqlConfig   *SqlConfig   `yaml:"mysql"`
	RedisConfig *RedisConfig `yaml:"redis"`
	JwtConfig   *JwtConfig   `yaml:"jwt"`
}

func init() {
	ReturnsInstance()
}

var Config = new(Info)
var filePath = "./conf/config.yaml"

type SqlConfig struct {
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Host     int    `yaml:"host"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type RedisConfig struct {
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

type JwtConfig struct {
	SigningKey  string `yaml:"signing-key"`
	ExpiresTime string `yaml:"expires-time"`
	BufferTime  string `yaml:"buffer-time"`
	Issuer      string `yaml:"issuer"`
}

func ReturnsInstance() *Info {
	Config = &Info{}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("读取配置文件失败,err : %s", err.Error())
	}
	err = yaml.Unmarshal(data, Config)
	if err != nil {
		log.Fatalf("解析配置文件失败,err : %s", err.Error())
	}
	fmt.Println(Config.JwtConfig)
	return Config
}
