package service

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"xhgblog/models"
	"xhgblog/utils/app"
	"xhgblog/utils/e"
)

type UserRegisterService struct {
	Email           string `form:"email" json:"email"`
	Password        string `form:"password" json:"password"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm"`
	//State           int    `form:"state" json:"state"`
}

func (this *UserRegisterService) UserRegValidation() *app.Response {
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
			Code: e.INVALID_PARAMS,
			Message:  msg,
		}
	}

	fmt.Println("Password", this.Password)
	fmt.Println("PasswordConfirm", this.PasswordConfirm)
	if this.Password != this.PasswordConfirm {
		return &app.Response{
			Code: e.ERROR_PASSWORD_DIFFER,
			Message:  e.GetMsg(e.ERROR_PASSWORD_DIFFER),
		}
	}

	return nil
}

func (this *UserRegisterService) UserRegister() (*models.User, *app.Response) {

	// 判断邮箱是否存在
	user, err := models.GetUserByEmail(this.Email)
	if err == nil && user.ID > 0 {
		return user, &app.Response{
			Code: e.ERROR_EXIST_EMAIL,
			Message:  e.GetMsg(e.ERROR_EXIST_EMAIL),
		}
	}

	user = &models.User{
		Email: this.Email,
		State: 0,
	}

	res := user.GeneratePassword(this.Password)
	if res != nil {
		return user, &app.Response{
			Code: e.ERROR_ENCRYPT,
			Message:  e.GetMsg(e.ERROR_ENCRYPT),
		}
	}

	err = models.AddUser(user)
	if err != nil {
		return user, &app.Response{
			Code: e.ERROR_CREATE_SQL,
			Message:  e.GetMsg(e.ERROR_CREATE_SQL),
		}
	}

	return user, nil
}
