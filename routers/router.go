package routers

import (
	"github.com/gin-gonic/gin"
	"xhgblog/controllers/user"
	"xhgblog/middleware"
	"xhgblog/utils/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Sessions(setting.AppSetting.Sessions.Secret))
	r.Use(middleware.CurrentUser())

	v1 := r.Group("app/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v1.POST("user/register", user.Register)
		v1.POST("user/login", user.Login)
		authed := v1.Group("/")
		authed.Use(middleware.AuthRequired())
		{
			authed.GET("user/me", user.UserMe)
			authed.DELETE("user/logout", user.Logout)
		}
	}
	return r
}
