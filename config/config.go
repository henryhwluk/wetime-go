package config

import (
	"fmt"
	"runtime"

	"github.com/jinzhu/configor"
)

type Config struct {
	AppName   string `default:"app name"`
	Server    string `default:"127.0.0.1"`
	HttpPort  string
	NginxPort string

	DB struct {
		Name     string `default:"wetime"`
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306"`
	}
	Contacts []struct {
		Name  string
		Email string `required:"true"`
	}
}

var Conf = Config{}

func Process() {
	sysType := runtime.GOOS
	var configPath string
	if sysType == "linux" {
		// LINUX系统
		configPath = "/opt/config.yml"
	}
	if sysType == "darwin" {
		// macOS
		configPath = "config.yml"
	}
	print(sysType)
	err := configor.Load(&Conf, configPath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v \n", Conf)
}
