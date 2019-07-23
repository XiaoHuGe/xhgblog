package admin

import (
	"github.com/Unknwon/com"
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

func DeleteTag(ctx *gin.Context) {
	G := app.Gin{C: ctx}
	resp := &app.Response{}

	id := com.StrTo(ctx.Param("id")).MustInt()
	tag, err := service.CheckTagByID(id)
	if err != nil {
		resp.Message = "没有此标签"
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}

	if tag.Total > 1 {
		editTagService := service.EditTagService{}
		editTagService.Total = tag.Total - 1
		editTagService.EdidTag(id)
		return
	}

	err = service.DeleteTag(id)
	if err != nil {
		resp.Message = "删除失败，内部错误"
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
