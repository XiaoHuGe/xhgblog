package routers

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
	"xhgblog/controllers"
	"xhgblog/controllers/admin"
	"xhgblog/controllers/home"
	"xhgblog/controllers/user"
	"xhgblog/middleware"
	"xhgblog/utils/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	setting.SetTemplate(r)
	gin.SetMode(setting.AppSetting.Server.RunMode)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Sessions(setting.AppSetting.Sessions.Secret))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	r.Static("/static", filepath.Join(setting.GetCurrentDirectory(), "./static"))
	r.NoRoute(controllers.NoRouterHtml)

	v1 := r.Group("/")
	{
		v1.GET("/", home.GetArticles)
		v1.GET("index", home.GetArticles)
		v1.GET("tag/:tag_id", home.GetArticlesByTag)
		v1.GET("archive/:year/:month", home.GetArticlesByArchive)
		v1.GET("category/:category_id", home.GetArticlesByCategory)
		v1.GET("article/:id", home.GetArticle)

		v1.GET("about", home.GetAbout)

		v1.GET("auth/:type", user.GetAuth)
		v1.GET("oauth/redirect", user.CallbackByAuth)
		v1.GET("captcha", home.GetCaptcha)

		v1.POST("visitor/comment", home.AddComment)

		us := v1.Group("/user")
		{
			if setting.AppSetting.Application.RegisterEnabled {
				us.GET("register", user.RegisterHtml)
				us.POST("register", user.Register)
			}
			us.GET("login", user.LoginHtml)
			us.POST("login", user.Login)
			us.GET("logout", user.Logout)
			us.GET("me", user.UserMe)
		}

		authed := v1.Group("/admin")
		authed.Use(middleware.AuthRequired())
		{
			authed.GET("index", admin.GetAdminIndex)

			//文章crud
			authed.GET("article", admin.GetArticles)
			authed.GET("new_article", admin.GetAddArticleHtml)
			authed.POST("new_article", admin.AddArticle)
			authed.POST("article/:id/delete", admin.DeleteArticle)
			authed.GET("article/:id/edit", admin.GetEditArticleHtml)
			authed.POST("article/:id/edit", admin.EditArticle)

			authed.GET("page", admin.GetPages)
			authed.GET("new_page", admin.GetAddPageHtml)
			authed.POST("new_page", admin.AddPage)
			authed.GET("page/:id/edit", admin.GetEditPageHtml)
			authed.POST("page/:id/edit", admin.EditPage)
			authed.POST("page/:id/delete", admin.DeletePage)

			authed.POST("tag/:id/delete", admin.DeleteTag)
			authed.POST("tag", admin.AddTag)

			authed.POST("category", admin.AddCategory)
		}
	}
	return r
}
