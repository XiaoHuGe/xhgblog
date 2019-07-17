package service

import "xhgblog/models"

type GetArticleService struct {
	TagID    int
	State    int
	PageNum  int
	PageSize int
}

func (this *GetArticleService) GetAll() ([]*models.Article, error) {
	maps := make(map[string]interface{})

	if this.State != -1 {
		maps["state"] = this.State
	}
	if this.TagID != -1 {
		maps["tag_id"] = this.TagID
	}
	return models.GetArticles(this.PageNum, this.PageSize, maps)
}

func (this *GetArticleService) Count() (int, error) {
	return models.GetArticleTotal(make(map[string]interface{}))
}

type AddArticleService struct {
	TagID     int    `form:"tag_id" json:"tag_id"`
	Title     string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Desc      string `form:"desc" json:"desc" binding:"required,min=2,max=100"`
	Content   string `form:"content" json:"content" binding:"required,min=2,max=3000"`
	CreatedBy string `form:"created_by" json:"created_by" binding:"required,min=2,max=30"`
}

func (this *AddArticleService) AddArticle() (error) {
	article := models.Article{
		TagID:     this.TagID,
		Title:     this.Title,
		Desc:      this.Desc,
		Content:   this.Content,
		CreatedBy: this.CreatedBy,
	}
	return models.AddArticle(&article)
}

type EditArcitleService struct {  // 长度验证问题
	TagID     int    `form:"tag_id" json:"tag_id"`
	Title     string `form:"title" json:"title"`
	Desc      string `form:"desc" json:"desc"`
	Content   string `form:"content" json:"content"`
	ModifiedBy string `form:"modified_by" json:"modified_by" binding:"required,min=2,max=30"`
}

func (this *EditArcitleService)EditArcitle(id int) error {
	article := models.Article{
		TagID:     this.TagID,
		Title:     this.Title,
		Desc:      this.Desc,
		Content:   this.Content,
		ModifiedBy: this.ModifiedBy,
	}
	return models.EditArticle(id, article)
}

func DeleteArticle(id int) error {
	return models.DeleteArticle(id)
}

func CheckArticleByID(id int) (bool, error) {
	return models.ExistArticleByID(id)
}