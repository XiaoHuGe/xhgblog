package service

import (
	"fmt"
	"strconv"
	"strings"
	"xhgblog/models"
)

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
	TagName     []string `form:"tag_name" json:"tag_name"`
	Tags        string   `form:"tags" json:"tags"`
	Title       string   `form:"title" json:"title" binding:"required,min=2,max=30"`
	Desc        string   `form:"desc" json:"desc"` // binding:"required,min=2,max=100"
	Content     string   `form:"content" json:"content" binding:"required,min=2,max=3000"`
	CreatedBy   string   `form:"created_by" json:"created_by"` // binding:"required,min=2,max=30"
	IsPublished string   `form:"is_published"`                 //json:"is_published"`
}

func (this *AddArticleService) AddArticle() (error) {
	tags := []models.Tag{}
	if len(this.Tags) > 0 {
		tagArr := strings.Split(this.Tags, ",")
		for _, tag := range tagArr {
			tagId, _ := strconv.ParseUint(tag, 10, 64)
			tag, _ := models.GetTag(int(tagId))
			tags = append(tags, tag)
		}
	}

	fmt.Println("isPublished :", this.IsPublished)
	published := "on" == this.IsPublished
	article := models.Article{
		//TagID:     this.TagID,
		Tags:        tags,
		Title:       this.Title,
		Desc:        this.Desc,
		Content:     this.Content,
		CreatedBy:   this.CreatedBy,
		IsPublished: published,
	}
	fmt.Println("add tags:", tags)
	return models.AddArticle(&article)
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

	article := models.Article{
		//TagID:     this.TagID,
		//Tags:       tags,  // update不更新tags
		Title:      this.Title,
		Desc:       this.Desc,
		Content:    this.Content,
		ModifiedBy: this.ModifiedBy,
	}
	err := models.EditArticle(id, &article)
	if err != nil {
		return err
	}
	// 此处逻辑待优化
	// 先删除文章下所有关联标签
	err = models.DeleteTagsByArticleId(id)
	if err != nil {
		return err
	}

	// 重新添加关联
	if len(this.Tags) > 0 {
		tagArr := strings.Split(this.Tags, ",")
		for _, tag := range tagArr {
			tagId, _ := strconv.ParseUint(tag, 10, 64)
			err := models.AddArticleJoinTags(id, int(tagId))
			if err != nil {
				return err
			}
		}
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
