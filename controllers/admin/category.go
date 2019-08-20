package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xhgblog/service"
	"xhgblog/utils/app"
)

func AddCategory(ctx *gin.Context) {
	G := app.Gin{C: ctx}
	resp := &app.Response{}

	addCategoryService := service.AddCategoryService{}
	err := ctx.ShouldBind(&addCategoryService)
	if err != nil {
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}
	c, err := addCategoryService.AddCategory()
	if err != nil {
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}
	resp.Code = 200
	resp.Message = "添加成功"
	resp.Data = c
	resp.Succeed = true
	G.Response(http.StatusOK, resp)
}
