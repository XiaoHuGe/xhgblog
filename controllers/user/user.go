package user

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"xhgblog/models"
	"xhgblog/service"
	"xhgblog/utils/app"
	"xhgblog/utils/e"
	"xhgblog/utils/setting"
)

func RegisterHtml(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "auth/signup.html", nil)
}

func LoginHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/signin.html", nil)
}

func Register(ctx *gin.Context) {
	G := app.Gin{C: ctx}

	registerService := service.UserRegisterService{}
	ctx.ShouldBind(&registerService)
	fmt.Println(registerService)
	resp := registerService.UserRegValidation()
	if resp != nil {
		G.Response(http.StatusOK, resp)
		return
	}

	user, resp := registerService.UserRegister()
	if resp != nil {
		G.Response(http.StatusOK, resp)
		return
	}

	resp = &app.Response{}
	data := service.BuildUserResponse(user)
	resp.Data = data
	resp.Code = e.SUCCESS_REGISTER
	resp.Message = e.GetMsg(resp.Code)
	resp.Succeed = true
	G.Response(http.StatusOK, resp)
}

func Login(ctx *gin.Context) {
	G := app.Gin{C: ctx}

	loginService := service.UserLoginService{}
	ctx.ShouldBind(&loginService)

	resp := loginService.UserLoginValidation()
	if resp != nil {
		G.Response(http.StatusOK, resp)
		return
	}

	user, resp := loginService.UserLogin()
	if resp != nil {
		G.Response(http.StatusOK, resp)
		return
	}

	// 设置session
	s := sessions.Default(ctx)
	s.Clear()
	s.Set(setting.SessionUserId, user.ID) // ctx.Set("user_id", user.ID)
	s.Save()
	ctx.Redirect(http.StatusMovedPermanently, "/admin/index")

	//resp = &app.Response{}
	//resp.Code = e.SUCCESS_LOGIN
	//resp.Message = e.GetMsg(resp.Code)
	//resp.Succeed = true
	//G.Response(http.StatusOK, resp)

}

func UserMe(ctx *gin.Context) {
	G := app.Gin{C: ctx}
	resp := &app.Response{}
	// ctx.Get("user")
	if u, _ := ctx.Get(setting.SessionUser); u != nil {
		if u, ok := u.(*models.User); ok {
			data := service.BuildUserResponse(u)
			resp.Data = data
			resp.Code = e.SUCCESS_GETME
			resp.Message = e.GetMsg(resp.Code)
		}
	}
	G.Response(http.StatusOK, resp)
}

func Logout(ctx *gin.Context) {
	//G := app.Gin{C: ctx}
	s := sessions.Default(ctx)
	s.Clear()
	s.Save()
	//resp := &app.Response{}
	//resp.Code = e.SUCCESS_LOGOUT
	//resp.Message = e.GetMsg(resp.Code)
	//G.Response(http.StatusOK, resp)
	ctx.Redirect(http.StatusSeeOther, "/user/login")

}
