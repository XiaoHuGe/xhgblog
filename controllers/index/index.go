package index

import (
	"fmt"
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
		//State:    state,
		PageNum:  util.GetPage(ctx),
		PageSize: setting.AppSetting.PageSize,
	}
	tags, err := getTagService.GetAll()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	totalPage := count / setting.AppSetting.PageSize
	if count % setting.AppSetting.PageSize > 0 {
		totalPage = count / setting.AppSetting.PageSize + 1
	}
	fmt.Printf("%d--%d--%d", count, totalPage, util.GetPage(ctx))
	ctx.HTML(http.StatusOK, "index/index.html", gin.H{
		"posts":           articles,
		"tags":            tags,
		"archives":       "",
		"links":           "",
		"user":            "",
		"pageIndex":       com.StrTo(ctx.Query("page")).MustInt(),
		"totalPage":       totalPage,
		"path":            ctx.Request.URL.Path,
		"maxReadPosts":    "",
		"maxCommentPosts": "",
	})
}
