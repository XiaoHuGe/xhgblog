package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xhgblog/service"
	"xhgblog/utils/app"
)

func AddTag(ctx *gin.Context) {
	G := app.Gin{C: ctx}
	resp := &app.Response{}

	addTagService := service.AddTagService{}
	err := ctx.ShouldBind(&addTagService)
	if err != nil {
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}

	tag, err := addTagService.AddTag()
	if err != nil {
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}
	resp.Code = 200
	resp.Message = "添加成功"
	resp.Data = tag
	resp.Succeed = true
	G.Response(http.StatusOK, resp)
}