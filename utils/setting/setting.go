package setting

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"xhgblog/utils/common"
)

const (
	SESSION_USER_ID      = "user_id"
	SESSION_USER         = "user"
	SESSION_GITHUB_STATE = "github_state"
	SESSION_CAPTCHA      = "session_captcha"
)

func SetTemplate(engine *gin.Engine) {

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
	engine.LoadHTMLGlob(filepath.Join(GetCurrentDirectory(), "./views/**/*"))
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
