package service

import (
	"fmt"
	"strings"
	"xhgblog/models"
)

type GetArticleService struct {
	TagID    int
	PageNum  int
	PageSize int
	Year     int
	Month    int
}

func (this *GetArticleService) GetAll() ([]*models.Article, error) {
	return models.GetArticles(this.TagID, this.PageNum, this.PageSize)
}

func (this *GetArticleService) GetArticlesByTagId() (articles []*models.Article, err error) {
	articles, err = models.GetArticlesByTagId(this.TagID, this.PageNum, this.PageSize)
	if err != nil {
		fmt.Println(err)
		return
	}
	return articles, nil
}

func (this *GetArticleService) GetArticlesByArchive() (articles []*models.Article, err error) {
	articles, err = models.GetArticlesByArchive(this.Year, this.Month, this.PageNum, this.PageSize)
	return articles, err
}

func (this *GetArticleService) Count() (int, error) {
	maps := make(map[string]interface{})

	if this.TagID != -1 {
		maps["tag_id"] = this.TagID
	}
	return models.GetArticleTotal(make(map[string]interface{}))
}

func (this *GetArticleService) GetCountByArchive() (int, error) {
	return models.GetArticleTotalByArchive(this.Year, this.Month)
}

func (this *GetArticleService) GetCountByTagId() (int, error) {
	return models.GetArticleTotalByTagId(this.TagID)
}

type AddArticleService struct {
	TagName     []string `form:"tag_name" json:"tag_name"`
	Tags        string   `form:"tags" json:"tags"`
	Title       string   `form:"title" json:"title" binding:"required,min=2,max=30"`
	Desc        string   `form:"desc" json:"desc"` // binding:"required,min=2,max=100"
	Content     string   `form:"content" json:"content" binding:"required,min=2,max=10000"`
	CreatedBy   string   `form:"created_by" json:"created_by"` // binding:"required,min=2,max=30"
	IsPublished string   `form:"is_published"`                 //json:"is_published"`
}

func (this *AddArticleService) AddArticle() error {
	//var tags []models.Tag
	var tagIds []string
	if len(this.Tags) > 0 {
		tagIds = strings.Split(this.Tags, ",")
	}
	tags, _ := models.GetTagsByIds(tagIds)

	//fmt.Println("isPublished :", this.IsPublished)
	published := "on" == this.IsPublished
	article := models.Article{
		//TagID:     this.TagID,
		Tags:        tags,
		Title:       this.Title,
		Content:     this.Content,
		CreatedBy:   this.CreatedBy,
		IsPublished: published,
	}
	//fmt.Println("add tags:", tags)
	err := models.AddArticle(&article)
	if err != nil {
		return err
	}
	return nil
}

type EditArcitleService struct {
	// 长度验证问题
	TagID       []int  `form:"tag_id" json:"tag_id"`
	Tags        string `form:"tags" json:"tags"`
	Title       string `form:"title" json:"title"`
	Desc        string `form:"desc" json:"desc"`
	Content     string `form:"content" json:"content"`
	ModifiedBy  string `form:"modified_by" json:"modified_by"` //binding:"required,min=2,max=30"
	IsPublished string `form:"is_published"`                   //json:"is_published"`
}

func (this *EditArcitleService) EditArcitle(id int) error {
	fmt.Println("EditArcitleService:", this.Tags)
	var tagIds []string
	if len(this.Tags) > 0 {
		tagIds = strings.Split(this.Tags, ",")
	}
	models.DeleteTagsByArticleId(id)
	tags, _ := models.GetTagsByIds(tagIds)
	article := models.Article{
		//TagID:     this.TagID,
		Tags:       tags,
		Title:      this.Title,
		Content:    this.Content,
		ModifiedBy: this.ModifiedBy,
	}
	err := models.EditArticle(id, &article)
	if err != nil {
		return err
	}
	return nil
}

func DeleteArticle(id int) error {
	return models.DeleteArticle(id)
}

func CheckArticleByID(id int) (bool, error) {
	return models.ExistArticleByID(id)
}

func GetArticle(id int) (*models.Article, error) {
	return models.GetArticle(id)
}
