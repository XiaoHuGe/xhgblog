package home

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
	"xhgblog/models"
	"xhgblog/service"
	"xhgblog/utils/app"
	"xhgblog/utils/setting"
	"xhgblog/utils/util"
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

	user, _ := ctx.Get(setting.SESSION_USER)

	ctx.HTML(http.StatusOK, "post/display.html", gin.H{
		"post": article,
		"user": user,
	})
}

func GetArticlesHtml(ctx *gin.Context) {
	// 文章名称
	var tagId int = -1
	if arg := ctx.Param("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
	}

	getArticleService := service.GetArticleService{
		TagID: tagId,
		//State:    state,
		PageNum:  util.GetPage(ctx),
		PageSize: setting.AppSetting.PageSize,
	}
	articles, err := getArticleService.GetAll()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	count, err := getArticleService.Count()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	getTagService := service.GetTagService{
		TagName:  "",
		PageNum:  util.GetPage(ctx),
		PageSize: setting.AppSetting.PageSize,
	}
	tags, err := getTagService.GetAll()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	pageIndex := com.StrTo(ctx.Query("page")).MustInt()
	if pageIndex < 1 {
		pageIndex = 1
	}
	totalPage := count / setting.AppSetting.PageSize
	if count%setting.AppSetting.PageSize > 0 {
		totalPage = count/setting.AppSetting.PageSize + 1
	}
	archives, err := models.GetArchive()
	user, _ := ctx.Get(setting.SESSION_USER)

	ctx.HTML(http.StatusOK, "index/index.html", gin.H{
		"posts":           articles,
		"tags":            tags,
		"archives":        archives,
		"user":            user,
		"pageIndex":       pageIndex,
		"totalPage":       totalPage,
		"path":            ctx.Request.URL.Path,
		"maxReadPosts":    "",
		"maxCommentPosts": "",
	})
}

func GetArticlesByTagHtml(ctx *gin.Context) {
	// 文章名称
	var tagId = -1
	if arg := ctx.Param("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
	}

	getArticleService := service.GetArticleService{
		TagID:    tagId,
		PageNum:  util.GetPage(ctx),
		PageSize: setting.AppSetting.PageSize,
	}
	articles, err := getArticleService.GetArticlesByTagId()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	count, _ := getArticleService.GetCountByTagId()

	getTagService := service.GetTagService{
		TagName:  "",
		PageNum:  util.GetPage(ctx),
		PageSize: setting.AppSetting.PageSize,
	}
	tags, err := getTagService.GetAll()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	pageIndex := com.StrTo(ctx.Query("page")).MustInt()
	if pageIndex < 1 {
		pageIndex = 1
	}
	totalPage := count / setting.AppSetting.PageSize
	if count%setting.AppSetting.PageSize > 0 {
		totalPage = count/setting.AppSetting.PageSize + 1
	}
	archives, err := models.GetArchive()
	user, _ := ctx.Get(setting.SESSION_USER)

	ctx.HTML(http.StatusOK, "index/index.html", gin.H{
		"posts":           articles,
		"tags":            tags,
		"archives":        archives,
		"user":            user,
		"pageIndex":       pageIndex,
		"totalPage":       totalPage,
		"path":            ctx.Request.URL.Path,
		"maxReadPosts":    "",
		"maxCommentPosts": "",
	})
}

func GetArticlesByArchiveHtml(ctx *gin.Context) {
	// 文章名称
	var tagId = -1
	var year int
	var month int
	if arg := ctx.Param("year"); arg != "" {
		year = com.StrTo(arg).MustInt()
	}
	if arg := ctx.Param("month"); arg != "" {
		month = com.StrTo(arg).MustInt()
	}

	getArticleService := service.GetArticleService{
		TagID:    tagId,
		Year:     year,
		Month:    month,
		PageNum:  util.GetPage(ctx),
		PageSize: setting.AppSetting.PageSize,
	}
	articles, err := getArticleService.GetArticlesByArchive()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	count, _ := getArticleService.GetCountByArchive()

	getTagService := service.GetTagService{
		TagName:  "",
		PageNum:  util.GetPage(ctx),
		PageSize: setting.AppSetting.PageSize,
	}
	tags, err := getTagService.GetAll()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	pageIndex := com.StrTo(ctx.Query("page")).MustInt()
	if pageIndex < 1 {
		pageIndex = 1
	}
	totalPage := count / setting.AppSetting.PageSize
	if count%setting.AppSetting.PageSize > 0 {
		totalPage = count/setting.AppSetting.PageSize + 1
	}
	archives, err := models.GetArchive()
	user, _ := ctx.Get(setting.SESSION_USER)

	ctx.HTML(http.StatusOK, "index/index.html", gin.H{
		"posts":           articles,
		"tags":            tags,
		"archives":        archives,
		"user":            user,
		"pageIndex":       pageIndex,
		"totalPage":       totalPage,
		"path":            ctx.Request.URL.Path,
		"maxReadPosts":    "",
		"maxCommentPosts": "",
	})
}
