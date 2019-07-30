package routers

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"path/filepath"
	"xhgblog/controllers"
	"xhgblog/controllers/admin"
	"xhgblog/controllers/home"
	"xhgblog/controllers/user"
	"xhgblog/middleware"
	"xhgblog/utils/common"
	"xhgblog/utils/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	setTemplate(r)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Sessions(setting.AppSetting.Sessions.Secret))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	r.Static("/static", filepath.Join("", "./static"))
	r.NoRoute(controllers.NoRouterHtml)

	v1 := r.Group("/")
	{
		v1.GET("/", home.GetArticlesHtml)
		v1.GET("index", home.GetArticlesHtml)
		v1.GET("tag/:tag_id", home.GetArticlesByTagHtml)
		v1.GET("archive/:year/:month", home.GetArticlesByArchiveHtml)
		v1.GET("article/:id", home.GetArticle)
		v1.GET("auth/:type", user.GetAuth)
		v1.GET("oauth/redirect", user.RedirectByAuth)
		v1.GET("captcha", home.GetCaptcha)
	}

	us := r.Group("/user")
	{
		//v1.GET("/", index.GetIndexHtml)
		if setting.AppSetting.Application.RegisterEnabled {
			us.GET("register", user.RegisterHtml)
			us.POST("register", user.Register)
		}
		us.GET("login", user.LoginHtml)
		us.POST("login", user.Login)
		us.GET("logout", user.Logout)
	}

	authed := v1.Group("/admin")
	authed.Use(middleware.AuthRequired())
	{
		authed.GET("index", admin.GetAdminIndexHtml)

		//文章crud
		authed.GET("article", admin.ManageArticleHtml)
		authed.GET("new_article", admin.GetAddArticleHtml)
		authed.POST("new_article", admin.AddArticle)
		authed.POST("article/:id/delete", admin.DeleteArticle)
		authed.GET("article/:id/edit", admin.GetEditArticleHtml)
		authed.POST("article/:id/edit", admin.EditArticle)

		authed.POST("tag/:id/delete", admin.DeleteTag)
		authed.POST("tag", admin.AddTag)

		authed.GET("user/me", user.UserMe)
		//authed.GET("user/logout", user.Logout)
	}
	return r
}

func setTemplate(engine *gin.Engine) {

	funcMap := template.FuncMap{
		"dateFormat": common.DateFormat,
		"substring":  common.Substring,
		"isOdd":      common.IsOdd,
		"isEven":     common.IsEven,
		"truncate":   common.Truncate,
		"add":        common.Add,
		"minus":      common.Minus,
		"listtag":    common.ListTag,
	}

	engine.SetFuncMap(funcMap)
	engine.LoadHTMLGlob(filepath.Join("", "./views/**/*"))
}
