package home

import (
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"xhgblog/models"
	"xhgblog/service"
	"xhgblog/utils/app"
	"xhgblog/utils/setting"
)

func AddComment(ctx *gin.Context) {
	G := app.Gin{C: ctx}
	resp := &app.Response{}

	addCommentService := service.AddCommentService{}
	err := ctx.ShouldBind(&addCommentService)
	if err != nil {
		resp.Message = "参数错误"
		resp.Error = err.Error()
		G.Response(http.StatusOK, resp)
		return
	}

	session := sessions.Default(ctx)
	cap := session.Get(setting.SESSION_CAPTCHA)
	session.Delete(setting.SESSION_CAPTCHA)
	if c, ok := cap.(string); ok {
		if !captcha.VerifyString(c, addCommentService.VerifyCode) {
			resp.Message = "验证码错误"
			G.Response(http.StatusOK, resp)
			return
		}
	}

	if user, _ := ctx.Get(setting.SESSION_USER); user != nil {
		u, ok := user.(*models.User);
		addCommentService.UserID = u.ID
		addCommentService.VisitorName = u.Email
		if ok && !u.IsAdmin {
			addCommentService.GithubUrl = u.GithubUrl
			addCommentService.AvatarUrl = u.AvatarUrl
			addCommentService.VisitorName = u.GithubLogin
		}
	}

	err = addCommentService.AddComment()
	if err != nil {
		resp.Message = "添加失败"
		G.Response(http.StatusOK, resp)
	}
	resp.Succeed = true
	G.Response(http.StatusOK, resp)
}
