package home

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
	"xhgblog/service"
	"xhgblog/utils/setting"
	"xhgblog/utils/util"
)

func GetIndexHtml(ctx *gin.Context) {
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
	totalPage := count / setting.AppSetting.PageSize
	if count%setting.AppSetting.PageSize > 0 {
		totalPage = count/setting.AppSetting.PageSize + 1
	}

	user, _ := ctx.Get(setting.SessionUser)

	ctx.HTML(http.StatusOK, "index/index.html", gin.H{
		"posts":           articles,
		"tags":            tags,
		"archives":        "",
		"links":           "",
		"user":            user,
		"pageIndex":       com.StrTo(ctx.Query("page")).MustInt(),
		"totalPage":       totalPage,
		"path":            ctx.Request.URL.Path,
		"maxReadPosts":    "",
		"maxCommentPosts": "",
	})
}

func GetIndexByTagHtml(ctx *gin.Context) {
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

	count := len(articles)

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
	totalPage := count / setting.AppSetting.PageSize
	if count%setting.AppSetting.PageSize > 0 {
		totalPage = count/setting.AppSetting.PageSize + 1
	}

	user, _ := ctx.Get(setting.SessionUser)

	ctx.HTML(http.StatusOK, "index/index.html", gin.H{
		"posts":           articles,
		"tags":            tags,
		"archives":        "",
		"links":           "",
		"user":            user,
		"pageIndex":       com.StrTo(ctx.Query("page")).MustInt(),
		"totalPage":       totalPage,
		"path":            ctx.Request.URL.Path,
		"maxReadPosts":    "",
		"maxCommentPosts": "",
	})
}
