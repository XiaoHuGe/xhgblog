package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"xhgblog/models"
	"xhgblog/utils/setting"
)

// 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		uid := session.Get(setting.SessionUserId)
		if uid != nil {
			user, err := models.GetUser(uid)
			if err == nil {
				ctx.Set(setting.SessionUser, &user)
			}
		}
		ctx.Next()
	}
}

// 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//G := app.Gin{C: ctx}

		if user, _ := ctx.Get(setting.SessionUser); user != nil {
			if _, ok := user.(*models.User); ok {
				ctx.Next()
				return
			}
		}
		//G.Response(http.StatusOK, &app.Response{
		//	Code:    e.ERROR_NOT_LOGIN,
		//	Message: e.GetMsg(e.ERROR_NOT_LOGIN),
		//})
		ctx.Redirect(http.StatusMovedPermanently, "/user/login")
		ctx.Abort()
	}
}
