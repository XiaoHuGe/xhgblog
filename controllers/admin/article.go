package admin

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
	"wblog/models"
	"xhgblog/service"
	"xhgblog/utils/app"
	"xhgblog/utils/setting"
	"xhgblog/utils/util"
)

func ManageArticleHtml(ctx *gin.Context) {
	G := app.Gin{C: ctx}
	resp := &app.Response{}
	// 文章名称
	var tagId = -1
	var state = -1

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
	user, _ := ctx.Get(setting.SessionUser)

	ctx.HTML(http.StatusOK, "admin/post.html", gin.H{
		"posts":  articles,
		"Active": "posts",
		"user":   user,
		//"comments": models.MustListUnreadComment(),
	})
}

func GetAddArticleHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "post/new.html", nil)
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
	ctx.Redirect(http.StatusMovedPermanently, "article")
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

func EditArticleHtml(ctx *gin.Context) {
	G := app.Gin{C: ctx}
	resp := &app.Response{}

	id := com.StrTo(ctx.Param("id")).MustInt()
	article, err := service.CheckArticleByID(id);
	if err != nil {
		resp.Message = "没有此文章" // 无效id
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}
	//article.Tags, _ = models.ListTagByPostId(id)  // 获取
	ctx.HTML(http.StatusOK, "post/modify.html", gin.H{
		"post": article,
	})
}

func EditArticle(ctx *gin.Context) {
	G := app.Gin{C: ctx}
	resp := &app.Response{}

	id := com.StrTo(ctx.Param("id")).MustInt()
	_, err := service.CheckArticleByID(id);
	if err != nil {
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
