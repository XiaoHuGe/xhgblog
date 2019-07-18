package models

import "github.com/jinzhu/gorm"

type Article struct {
	Model
	TagID int   `json:"tag_id" gorm:"index"`
	Tags  []Tag `json:"tags" gorm:"many2many:article_tag;"`

	Title       string `json:"title"`
	Desc        string `json:"desc"`
	Content     string `json:"content"`
	IsPublished bool   `json:"is_published"` // published or not
	CreatedBy   string `json:"created_by"`
	ModifiedBy  string `json:"modified_by"`
}

func GetArticles(pageNum int, pageSize int, maps interface{}) ([]*Article, error) {
	var articles []*Article
	//var tags []Tag
	err := db.Preload("Tags").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}

// 获取单个文章
func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	err = db.Model(&article).Related(&article.Tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &article, nil
}

func DeleteArticle(id int) (error) {
	err := db.Where("id = ?", id).Delete(&Article{}).Error
	if err != nil {
		return err
	}
	return nil
}

func EditArticle(id int, article *Article) (error) {
	err := db.Model(&Article{}).Where("id = ? ", id).Update(article).Error
	if err != nil {
		return err
	}
	return nil
}

func AddArticle(article *Article) (error) {
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
	err := db.Where("id = ? ", id).First(&Article{}).Error
	if err != nil && err == gorm.ErrRecordNotFound { // 错误不为空且为未找到时返回false
		return false, err
	}
	return true, nil
}
