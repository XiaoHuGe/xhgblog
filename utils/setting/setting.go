package setting

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

var (
	SessionUserId = "user_id"
	SessionUser   = "user"
)

type App struct {
	Server
	Database
	Sessions
}

var AppSetting = &App{}

type Server struct {
	RunMode      string        `yaml:"RunMode"`
	HttpPort     int           `yaml:"HttpPort"`
	ReadTimeout  time.Duration `yaml:"ReadTimeoutmeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
}

type Database struct {
	Type        string `yaml:"Type"`
	User        string `yaml:"User"`
	Password    string `yaml:"Password"`
	Host        string `yaml:"Host"`
	Name        string `yaml:"Name"`
	TablePrefix string `yaml:"TablePrefix"`
}

type Sessions struct {
	Secret string `yaml:"Secret"`
}

func Setup() {
	yamlFile, err := ioutil.ReadFile("conf/app.yaml")
	if err != nil {
		fmt.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, AppSetting)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}
	AppSetting.Server.ReadTimeout = AppSetting.Server.ReadTimeout * time.Second
	AppSetting.Server.WriteTimeout = AppSetting.Server.WriteTimeout * time.Second
	fmt.Printf("yaml %v", AppSetting.Server.HttpPort)
}
