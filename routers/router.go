package routers

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"path/filepath"
	"xhgblog/controllers/admin"
	"xhgblog/controllers/article"
	"xhgblog/controllers/index"
	"xhgblog/controllers/tag"
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

	v1 := r.Group("/")
	{
		v1.GET("/", index.IndexHtml)
		v1.GET("register", user.RegisterHtml)
		v1.POST("register", user.Register)
		v1.GET("login", user.LoginHtml)
		v1.POST("login", user.Login)
		authed := v1.Group("/admin")
		authed.Use(middleware.AuthRequired())
		{
			authed.GET("index", admin.AdminIndex)
			authed.GET("user/me", user.UserMe)
			authed.DELETE("user/logout", user.Logout)
		}

		// 获取标签列表
		v1.GET("/tags", tag.GetTags)
		// 新建标签
		v1.POST("/tag", tag.AddTag)
		// 修改标签
		v1.PUT("/tag/:id", tag.EditTag)
		// 删除标签
		v1.DELETE("/tag/:id", tag.DeleteTag)

		// 获取文章列表
		v1.GET("/articles", article.GetArticles)
		// 获取指定文章
		v1.GET("/article/:id", article.GetArticle)
		// 新建文章
		v1.POST("/article", article.AddArticle)
		// 修改文章
		v1.PUT("/article/:id", article.EditArticle)
		// 删除文章
		v1.DELETE("/article/:id", article.DeleteArticle)

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