package home

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
	"xhgblog/service"
	"xhgblog/utils/app"
	"xhgblog/utils/setting"
)

func GetArticle(ctx *gin.Context) {
	G := app.Gin{C: ctx}
	resp := &app.Response{}

	id := com.StrTo(ctx.Param("id")).MustInt()
	b, err := service.CheckArticleByID(id)
	if b == false {
		resp.Message = "没有此文章" // 无效id
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}
	article, err := service.GetArticle(id)
	if err != nil {
		resp.Message = "获取文章失败"
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}

	user, _ := ctx.Get(setting.SessionUser)

	ctx.HTML(http.StatusOK, "post/display.html", gin.H{
		"post": article,
		"user": user,
	})
}
