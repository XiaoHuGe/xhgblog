package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Model
	//TagID int   `json:"tag_id" gorm:"index"`
	Tags        []Tag  `json:"tags" gorm:"many2many:article_tags;"` //table article_tags
	Title       string `json:"title"`
	Content     string `json:"content" gorm:"size:3000"`
	IsPublished bool   `json:"is_published"` // published or not
	CreatedBy   string `json:"created_by"`
	ModifiedBy  string `json:"modified_by"`
}

func GetArticles(tagId, pageNum, pageSize int) ([]*Article, error) {
	var articles []*Article
	err := db.Preload("Tags").Offset(pageNum).Limit(pageSize).Order("created_at desc").Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}

//从归档获取文章列表
func GetArticlesByArchive(year, month, pageNum, pageSize int) ([]*Article, error) {
	var articles []*Article
	date := fmt.Sprintf("%d-%02d", year, month)
	err := db.Preload("Tags").Where("DATE_FORMAT(created_at,'%Y-%m') = ?", date).Offset(pageNum).Limit(pageSize).Order("created_at desc").Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}

func GetArticleTotalByArchive(year, month int) (int, error) {
	var articles []*Article
	var count int
	date := fmt.Sprintf("%d-%02d", year, month)
	err := db.Where("DATE_FORMAT(created_at,'%Y-%m') = ?", date).Find(&articles).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetArticlesByTagId(tagId, pageNum, pageSize int) ([]*Article, error) {
	var articles []*Article
	//err := db.Preload("Tags").Where("DATE_FORMAT(created_at,'%Y-%m') = ?", date).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	err := db.Preload("Tags").
		Joins("JOIN xhgblog_article_tags ON xhgblog_article_tags.article_id = xhgblog_article.id").
		Joins("JOIN xhgblog_tag ON xhgblog_tag.id = xhgblog_article_tags.tag_id").
		Where("xhgblog_article_tags.tag_id = ?", tagId).
		Offset(pageNum).Limit(pageSize).Order("created_at desc").
		Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}

// 获取单个文章
func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Preload("Tags").Where("id = ? ", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &article, nil
}

func DeleteArticle(id int) error {
	err := db.Where("id = ?", id).Delete(&Article{}).Error
	if err != nil {
		return err
	}
	return nil
}

func EditArticle(id int, article *Article) error {
	var arti Article
	//err := db.Model(&arti).Where("id = ? ", id).Update(article).Error
	err := db.Preload("Tags").Where("id = ? ", id).First(&arti).Update(article).Error
	if err != nil {
		return err
	}
	return nil
}

func AddArticle(article *Article) error {
	err := db.Create(article).Error
	if err != nil {
		return err
	}
	return nil
}

func GetArticleTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Article{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := db.Where("id = ? ", id).First(&article).Error
	if err != nil && err == gorm.ErrRecordNotFound { // 错误不为空且为未找到时返回false
		return false, err
	}
	return true, nil
}
