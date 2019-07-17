package service

import (
	"time"
	"xhgblog/models"
)

// 返回信息
type UserRegisterResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	State     int       `json:"state"`
	CreatedAt time.Time `json:"created_at"`
}

func BuildUserResponse(user *models.User) *UserRegisterResponse {
	return &UserRegisterResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		State:     user.State,
	}
}
