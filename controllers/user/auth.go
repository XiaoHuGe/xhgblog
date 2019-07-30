package user

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/snluu/uuid"
	"net/http"
	"xhgblog/utils/setting"
)

func GetAuth(c *gin.Context) {
	authType := c.Param("type")

	session := sessions.Default(c)
	uuid := uuid.Rand().Hex()
	session.Delete(setting.SESSION_GITHUB_STATE)
	session.Set(setting.SESSION_GITHUB_STATE, uuid)
	session.Save()

	authurl := "/user/login"
	switch authType {
	case "github":
		authurl = fmt.Sprintf(setting.AppSetting.OAuth.GithubAuthUrl, setting.AppSetting.OAuth.GithubClientID, uuid)
	case "weibo":
	default:
	}
	c.Redirect(http.StatusFound, authurl)
}