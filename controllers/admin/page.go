package admin

import (
	"fmt"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
	"xhgblog/models"
	"xhgblog/service"
	"xhgblog/utils/app"
	"xhgblog/utils/setting"
)

func ManagePageHtml(ctx *gin.Context) {
	pages, _ := models.GetPages()
	user, _ := ctx.Get(setting.SESSION_USER)
	ctx.HTML(http.StatusOK, "admin/page.html", gin.H{
		"pages":    pages,
		"user":     user,
		"comments": nil,
	})
}

func GetAddPageHtml(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "page/new.html", nil)
}

func AddPage(ctx *gin.Context) {
	G := app.Gin{C: ctx}
	resp := &app.Response{}

	pageService := service.PageService{}
	err := ctx.ShouldBind(&pageService)
	if err != nil {
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
	}

	err = pageService.AddPage()
	if err != nil {
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
	}
	ctx.Redirect(http.StatusMovedPermanently, "/admin/page")
}

func GetEditPageHtml(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()
	pageService := service.PageService{}
	page, _ := pageService.GetPageByID(id)
	ctx.HTML(http.StatusOK, "page/modify.html", gin.H{
		"page": page,
	})
}

func EditPage(ctx *gin.Context) {
	G := app.Gin{C: ctx}
	resp := &app.Response{}

	id := com.StrTo(ctx.Param("id")).MustInt()
	fmt.Println("id : ", id)
	pageService := service.PageService{}
	err := ctx.ShouldBind(&pageService)
	if err != nil {
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
	}

	err = pageService.EditPage(id)
	if err != nil {
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
	}
	ctx.Redirect(http.StatusMovedPermanently, "/admin/page")
}

func DeletePage(ctx *gin.Context) {
	G := app.Gin{C: ctx}
	resp := &app.Response{}

	id := com.StrTo(ctx.Param("id")).MustInt()
	pageService := service.PageService{}
	err := ctx.ShouldBind(&pageService)
	if err != nil {
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
	}
	err = pageService.DeletePageByID(id)
	if err != nil {
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
	}
	resp.Succeed = true
	resp.Message = "删除成功"
	G.Response(http.StatusOK, resp)
}