package service

import (
	"github.com/astaxie/beego/validation"
	"xhgblog/models"
	"xhgblog/utils/app"
	"xhgblog/utils/e"
)

type UserLoginService struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (this *UserLoginService) UserLoginValidation() *app.Response {
	valid := validation.Validation{}
	valid.Required(this.Email, "email").Message("邮箱不能为空")
	valid.MaxSize(this.Email, 20, "email").Message("邮箱长度不能超过20")
	valid.Required(this.Password, "password").Message("密码不能为空")
	valid.MaxSize(this.Password, 16, "passwordMax").Message("密码长度不能大于16个字符")
	valid.MinSize(this.Password, 6, "passwordMin").Message("密码长度不能少于6个字符")

	if valid.HasErrors() {
		msg := make([]string, len(valid.Errors))
		for i, err := range valid.Errors {
			msg[i] = err.Message
		}
		return &app.Response{
			Code:    e.INVALID_PARAMS,
			Message: msg,
		}
	}

	return nil
}

func (this *UserLoginService) UserLogin() (models.User, *app.Response, error) {
	//var user models.User
	//user.State = 1
	user, err := models.IsExistUserByEmail(this.Email)
	if err != nil {
		return *user, &app.Response{
			Code:    e.ERROR_NOT_ENAIL,
			Message: e.GetMsg(e.ERROR_NOT_ENAIL),
		}, err
	}

	//err = bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(this.Password))
	err = user.CheckPassword(this.Password)
	if err != nil {
		return *user, &app.Response{
			Code:    e.ERROR_ENAIL_OR_PASS,
			Message: e.GetMsg(e.ERROR_ENAIL_OR_PASS),
		}, err
	}
	return *user, nil, nil
}
