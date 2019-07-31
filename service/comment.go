package service

import (
	"xhgblog/models"
)

type AddCommentService struct {
	Content   string `form:"content"`
	UserID    uint
	ArticleID uint `form:"article_id"`

	AvatarUrl   string
	VisitorName string
	GithubUrl   string
	VerifyCode  string `form:"verify_code"`
}

func (this *AddCommentService) AddComment() error {

	comment := &models.Comment{
		Content:     this.Content,
		UserID:      this.UserID,
		ArticleID:   this.ArticleID,
		AvatarUrl:   this.AvatarUrl,
		VisitorName: this.VisitorName,
		GithubUrl:   this.GithubUrl,
	}
	err := models.AddComment(comment)
	if err != nil {
		return err
	}
	return nil
}
