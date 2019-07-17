package tag

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
	"xhgblog/service"
	"xhgblog/utils/app"
	//"xhgblog/utils/e"
	"xhgblog/utils/setting"
	"xhgblog/utils/util"
)

type respData struct {
	Lists interface{} `json:"lists"`
	Total int         `json:"total"`
}

func GetTags(ctx *gin.Context) {
	G := &app.Gin{C: ctx}
	resp := &app.Response{}

	// 标签名称
	name := ctx.Query("name")

	var state int = -1
	if arg := ctx.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	getTagService := service.GetTagService{
		TagName:  name,
		State:    state,
		PageNum:  util.GetPage(ctx),
		PageSize: setting.AppSetting.PageSize,
	}
	tags, err := getTagService.GetAll()
	if err != nil {
		resp.Message = "获取标签列表失败"
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}

	count, err := getTagService.Count()
	if err != nil {
		resp.Message = "获取标签数量失败"
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}

	rd := respData{}
	rd.Lists = tags
	rd.Total = count
	resp.Data = rd
	resp.Succeed = true
	G.Response(http.StatusOK, resp)
}

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
	err = addTagService.AddTag()
	if err != nil {
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}

	resp.Code = 200
	resp.Message = "添加成功"
	resp.Succeed = true
	G.Response(http.StatusOK, resp)
}

func EditTag(ctx *gin.Context) {
	G := app.Gin{C: ctx}
	resp := &app.Response{}
	
	id := com.StrTo(ctx.Param("id")).MustInt()
	b, err := service.CheckID(id);
	if b == false{
		resp.Message = "没有此标签"  // 无效id
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}
	editTagService := service.EditTagService{}
	err = ctx.ShouldBind(&editTagService)
	if err != nil {
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}
	
	err = editTagService.EdidTag(id)
	if err != nil {
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}

	resp.Succeed = true
	resp.Message = "修改成功"
	G.Response(http.StatusOK, resp)
}

func DeleteTag(ctx *gin.Context) {
	G := app.Gin{C: ctx}
	resp := &app.Response{}

	id := com.StrTo(ctx.Param("id")).MustInt()
	b, err := service.CheckID(id);
	if b == false{
		resp.Message = "没有此标签"
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}

	err = service.DeleteTag(id)
	if err != nil {
		resp.Message = "删除失败，内部错误"
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}

	resp.Succeed = true
	resp.Message = "删除成功"
	G.Response(http.StatusOK, resp)
}
