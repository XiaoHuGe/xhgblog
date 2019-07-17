package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"xhgblog/models"
	"xhgblog/utils/app"
	"xhgblog/utils/e"
	"xhgblog/utils/setting"
)

// 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(cxt *gin.Context) {
		session := sessions.Default(cxt)
		uid := session.Get(setting.SessionUserId)
		if uid != nil {
			user, err := models.GetUser(uid)
			if err == nil {
				cxt.Set(setting.SessionUser, &user)
			}
		}
		cxt.Next()
	}
}

// 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(cxt *gin.Context) {
		G := app.Gin{C: cxt}

		if user, _ := cxt.Get(setting.SessionUser); user != nil {
			if _, ok := user.(*models.User); ok {
				cxt.Next()
				return
			}
		}
		G.Response(http.StatusOK, &app.Response{
			Code: e.ERROR_NOT_LOGIN,
			Message:  e.GetMsg(e.ERROR_NOT_LOGIN),
		})
		cxt.Abort()
	}
}
