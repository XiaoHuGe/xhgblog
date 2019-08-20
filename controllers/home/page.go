package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xhgblog/service"
	"xhgblog/utils/app"
	"xhgblog/utils/setting"
)

func GetAbout(ctx *gin.Context) {
	G := app.Gin{C: ctx}
	resp := &app.Response{}

	pageService := service.PageService{}
	page, err := pageService.GetPageByTitle("关于")
	if err != nil {
		resp.Message = "获取内容失败"
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}
	user, _ := ctx.Get(setting.SESSION_USER)

	ctx.HTML(http.StatusOK, "page/display.html", gin.H{
		"page": page,
		"user": user,
	})
}
