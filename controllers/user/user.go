package user

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"xhgblog/models"
	"xhgblog/service"
	"xhgblog/utils/app"
	"xhgblog/utils/e"
	"xhgblog/utils/log"
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
	//G := app.Gin{C: ctx}

	loginService := service.UserLoginService{}
	ctx.ShouldBind(&loginService)

	resp := loginService.UserLoginValidation()
	if resp != nil {
		//G.Response(http.StatusOK, resp)
		ctx.HTML(http.StatusForbidden, "errors/error.html", gin.H{
			"message": resp.Message,
		})
		return
	}

	user, resp, err := loginService.UserLogin()
	if err != nil {
		//G.Response(http.StatusOK, resp)
		ctx.HTML(http.StatusForbidden, "errors/error.html", gin.H{
			"message": resp.Message,
		})
		return
	}

	// 设置session
	s := sessions.Default(ctx)
	s.Clear()
	s.Set(setting.SESSION_USER_ID, user.ID) // ctx.Set("user_id", user.ID)
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
	if u, _ := ctx.Get(setting.SESSION_USER); u != nil {
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

type GithubUserInfo struct {
	AvatarURL string `json:"avatar_url"`
	Login     string `json:"login"`
	HTMLURL   string `json:"html_url"`
}

func CallbackByAuth(ctx *gin.Context) {

	code := ctx.Query("code")
	state := ctx.Query("state")
	session := sessions.Default(ctx)
	if len(state) == 0 || session.Get(setting.SESSION_GITHUB_STATE) != state {
		fmt.Println("session.Get err:", session.Get(setting.SESSION_GITHUB_STATE))
		ctx.Abort()
		return
	}
	session.Delete(setting.SESSION_GITHUB_STATE)
	session.Save()

	token, err := GetTokenByCode(ctx, code)
	if err != nil {
		fmt.Println("GetTokenByCode err :", err)
		return
	}
	userInfo, err := GetUserIndoByToken(token)
	user := &models.User{
		GithubLogin: userInfo.Login,
		GithubUrl:   userInfo.HTMLURL,
		AvatarUrl:   userInfo.AvatarURL,
		State:       0,
	}
	err = models.FirstOrAddUser(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Logrus.Info("github user login, github_login:", userInfo.Login)
	// 设置session
	s := sessions.Default(ctx)
	s.Clear()
	s.Set(setting.SESSION_USER_ID, user.ID) // ctx.Set("user_id", user.ID)
	s.Save()
	ctx.Redirect(http.StatusMovedPermanently, "/")
	return

}

func GetTokenByCode(ctx *gin.Context, code string) (string, error) {
	conf := &oauth2.Config{
		ClientID:     setting.AppSetting.OAuth.GithubClientID,
		ClientSecret: setting.AppSetting.OAuth.GithubClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  setting.AppSetting.OAuth.GithubAuthUrl,
			TokenURL: setting.AppSetting.OAuth.GithubTokenUrl,
		},
	}
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		return "", err
	}
	token := tok.AccessToken
	return token, nil
}

func GetUserIndoByToken(token string) (*GithubUserInfo, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/user?access_token=%s", token))
	if err != nil {
		fmt.Println("http.Get err :", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll err :", err)
		return nil, err
	}
	var userInfo GithubUserInfo
	err = json.Unmarshal(body, &userInfo)
	return &userInfo, nil
}
