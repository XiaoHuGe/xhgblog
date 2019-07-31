package models

type Comment struct {
	Model
	Content   string `json:"content"`
	UserID    uint   `json:"user_id"`
	ArticleID uint   `json:"article_id"`

	AvatarUrl   string `json:"avatar_url"`
	VisitorName string `json:"visitor_name"`
	GithubUrl   string `json:"github_url"`
}

func AddComment(comment *Comment) error {
	return db.Create(comment).Error
}