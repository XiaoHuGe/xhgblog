package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"syscall"
	"xhgblog/models"
	"xhgblog/routers"
	"xhgblog/utils/log"
	"xhgblog/utils/setting"
)

func main() {
	setting.Setup()
	log.SetUp()
	models.Setup()

	r := routers.InitRouter()

	readTimeout := setting.AppSetting.Server.ReadTimeout
	writeTimeout := setting.AppSetting.Server.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.AppSetting.Server.HttpPort)
	maxHeaderBytes := 1 << 20

	endless.DefaultReadTimeOut = readTimeout
	endless.DefaultWriteTimeOut = writeTimeout
	endless.DefaultMaxHeaderBytes = maxHeaderBytes
	server := endless.NewServer(endPoint, r)
	server.BeforeBegin = func(add string) {
		fmt.Printf("Actual pid is %d\n", syscall.Getpid())
		log.Logrus.Info("Actual pid is ", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		//fmt.Printf("Server err: %v\n", err)
		log.Logrus.Error("Server err: %v\n", err)
		return
	}
	//r.Run(fmt.Sprintf(":%d", setting.AppSetting.HttpPort))
}
