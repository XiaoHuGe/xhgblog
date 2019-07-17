package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xhgblog/utils/setting"
)

func AdminIndex(ctx *gin.Context) {
	user, _ := ctx.Get(setting.SessionUser)
	ctx.HTML(http.StatusOK, "admin/index.html", gin.H{
		"pageCount":    5,
		"postCount":    5,
		"tagCount":     5,
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
