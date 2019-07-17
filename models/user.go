package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model
	Email          string `json:"email"`
	Username       string `json:"username"`
	PasswordDigest string `json:"password_digest"`
	State          int    `json:"state"`
}

func GetUser(ID interface{}) (User, error)  {
	var user User
	result := db.First(&user, ID)
	return user, result.Error
}

func AddUser(user *User) error {
	return db.Create(user).Error
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &user, err
	}
	return &user, nil
}

func DeleteUser(user *User) error {
	return db.Delete(user).Error
}

// 生成密码
func (this *User)GeneratePassword(password string) error {
	data, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	this.PasswordDigest = string(data)
	return nil
}

// 验证密码
func (this *User)CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(this.PasswordDigest), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
