package main

import (
	"fmt"
	"xhgblog/models"
	"xhgblog/routers"
	"xhgblog/utils/setting"
)

func main() {
	setting.Setup()
	models.Setup()

	r := routers.InitRouter()

	r.Run(fmt.Sprintf(":%d", setting.AppSetting.HttpPort))
}
