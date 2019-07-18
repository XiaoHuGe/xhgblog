package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xhgblog/service"
	"xhgblog/utils/app"
	"xhgblog/utils/setting"
	"xhgblog/utils/util"
)

func GetArticleHtml(ctx *gin.Context) {
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
		"posts":    articles,
		"Active":   "posts",
		"user":     user,
		//"comments": models.MustListUnreadComment(),
	})
}

func AddArticleHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "post/new.html", nil)
}