package setting

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type App struct {
	Application
	Server
	Database
	Sessions
	OAuth
}

var AppSetting = &App{}

type Application struct {
	PageSize        int    `yaml:"PageSize"`
	JwtSecret       string `yaml:"JwtSecret"`
	RegisterEnabled bool   `yaml:"RegisterEnabled"`
}

type Server struct {
	RunMode      string        `yaml:"RunMode"`
	HttpPort     int           `yaml:"HttpPort"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
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

type OAuth struct {
	GithubClientID     string `yaml:"GithubClientID"`
	GithubClientSecret string `yaml:"GithubClientSecret"`
	GithubRedirectUrl  string `yaml:"GithubRedirectUrl"`
	GithubAuthUrl      string `yaml:"GithubAuthUrl"`
	GithubTokenUrl     string `yaml:"GithubTokenUrl"`
}

func Setup() {

	yamlFile, err := ioutil.ReadFile("conf/app.yaml")
	if err != nil {
		//fmt.Printf("yamlFile.Get err   #%v ", err)
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, AppSetting)
	if err != nil {
		//fmt.Printf("Unmarshal: %v", err)
		panic(err)
	}
	AppSetting.Server.ReadTimeout = AppSetting.Server.ReadTimeout * time.Second
	AppSetting.Server.WriteTimeout = AppSetting.Server.WriteTimeout * time.Second
	fmt.Printf("yaml %v", AppSetting.Server.HttpPort)
}
