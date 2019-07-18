package article

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
	"xhgblog/service"
	"xhgblog/utils/app"
	"xhgblog/utils/setting"
	"xhgblog/utils/util"
)

type respData struct {
	Lists interface{} `json:"lists"`
	Total int         `json:"total"`
}

func GetArticles(ctx *gin.Context) {
	G := &app.Gin{C: ctx}
	resp := &app.Response{}

	// 文章名称
	var tagId int = -1
	if arg := ctx.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
	}

	var state int = -1
	if arg := ctx.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	getArticleService := service.GetArticleService{
		TagID:    tagId,
		State:    state,
		PageNum:  util.GetPage(ctx),
		PageSize: setting.AppSetting.PageSize,
	}
	articles, err := getArticleService.GetAll()
	if err != nil {
		resp.Message = "获取文章列表失败"
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}

	count, err := getArticleService.Count()
	if err != nil {
		resp.Message = "获取文章数量失败"
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}

	rd := respData{}
	rd.Lists = articles
	rd.Total = count
	resp.Data = rd
	resp.Succeed = true
	G.Response(http.StatusOK, resp)
}

func GetArticle(ctx *gin.Context) {
	G := app.Gin{C: ctx}
	resp := &app.Response{}

	id := com.StrTo(ctx.Param("id")).MustInt()
	b, err := service.CheckArticleByID(id);
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
	//resp.Data = article
	//resp.Succeed = true
	//G.Response(http.StatusOK, resp)
	//article.Tag
	ctx.HTML(http.StatusOK, "post/display.html", gin.H{
		"post": article,
		//"user": user,
	})
}

func AddArticle(ctx *gin.Context) {
	G := app.Gin{C: ctx}
	resp := &app.Response{}

	addArticleService := service.AddArticleService{}
	err := ctx.ShouldBind(&addArticleService)
	if err != nil {
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}
	err = addArticleService.AddArticle()
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

func EditArticle(ctx *gin.Context) {
	G := app.Gin{C: ctx}
	resp := &app.Response{}

	id := com.StrTo(ctx.Param("id")).MustInt()
	b, err := service.CheckArticleByID(id);
	if b == false {
		resp.Message = "没有此文章" // 无效id
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}
	editArcitleService := service.EditArcitleService{}
	err = ctx.ShouldBind(&editArcitleService)
	if err != nil {
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}

	err = editArcitleService.EditArcitle(id)
	if err != nil {
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}

	resp.Succeed = true
	resp.Message = "修改成功"
	G.Response(http.StatusOK, resp)
}

func DeleteArticle(ctx *gin.Context) {
	G := app.Gin{C: ctx}
	resp := &app.Response{}

	id := com.StrTo(ctx.Param("id")).MustInt()
	b, err := service.CheckArticleByID(id);
	if b == false {
		resp.Message = "没有此文章"
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}

	err = service.DeleteArticle(id)
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
