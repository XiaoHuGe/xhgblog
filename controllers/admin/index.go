package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xhgblog/service"
	"xhgblog/utils/setting"
)

func GetAdminIndexHtml(ctx *gin.Context) {
	user, _ := ctx.Get(setting.SessionUser)

	getArticleService := &service.GetArticleService{}
	postCount, _ := getArticleService.Count()

	getTagService := &service.GetTagService{}
	tagCount, _ := getTagService.Count()

	fmt.Printf("postCount=%d, tagCount=%d \n", postCount, tagCount)

	ctx.HTML(http.StatusOK, "admin/index.html", gin.H{
		"pageCount":    5,
		"postCount":    postCount,
		"tagCount":     tagCount,
		"commentCount": 5,
		"user":         user,
	})
}

//user, _ := ctx.Get(setting.SessionUser)
//ctx.HTML(http.StatusOK, "admin/index.html", gin.H{
//"pageCount":    models.CountPage(),
//"postCount":    models.CountPost(),
//"tagCount":     models.CountTag(),
//"commentCount": models.CountComment(),
//"user":         user,
//"comments":     models.MustListUnreadComment(),
//})
