package home

import (
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"xhgblog/utils/setting"
)

func GetCaptcha(ctx *gin.Context) {
	session := sessions.Default(ctx)
	cap := captcha.New()
	session.Delete(setting.SESSION_CAPTCHA)
	session.Set(setting.SESSION_CAPTCHA, cap)
	session.Save()
	captcha.WriteImage(ctx.Writer, cap, 120, 40)
}
