package models

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model
	TagName string //`json:"tag_name"`
	Total   int    `json:"total" gorm:"-"` // article num
}

func GetTags() ([]*Tag, error) {
	var (
		tags []*Tag
		err  error
	)
	rows, err := db.Raw("select t.*,count(*) total from xhgblog_tag t inner join xhgblog_article_tags pt on t.id = pt.tag_id inner join xhgblog_article p on pt.article_id = p.id group by pt.tag_id").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tag Tag
		db.ScanRows(rows, &tag)
		tags = append(tags, &tag)
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tags, err
}

func GetTag(id int) (Tag, error) {
	var tag Tag
	err := db.Where("id = ? ", id).Find(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tag, err
	}
	return tag, nil
}

func GetTagsByIds(ids []string) ([]Tag, error) {
	var tags []Tag
	err := db.Where("id IN (?)", ids).Find(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tags, err
	}
	return tags, nil
}

func GetTagByName(name string) (Tag, error) {
	var tag Tag
	err := db.Where("tag_name = ? ", name).Find(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tag, err
	}
	return tag, nil
}

func DeleteTag(id int) error {
	err := db.Where("id = ?", id).Delete(&Tag{}).Error
	if err != nil {
		return err
	}
	return nil
}

func EditTag(id int, data interface{}) error {
	err := db.Model(&Tag{}).Where("id = ?", id).Update(data).Error
	if err != nil {
		return err
	}
	return nil
}

func AddTag(name string) (*Tag, error) {
	var tag Tag
	err := db.Where(Tag{TagName: name}).FirstOrCreate(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &tag, nil
}

func GetTagTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Tag{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func ExistTagByID(id int) (*Tag, error) {
	var tag Tag
	err := db.Where("id = ? ", id).First(&tag).Error
	if err != nil && err == gorm.ErrRecordNotFound { // 错误不为空且为未找到时返回false
		return nil, err
	}
	return &tag, nil
}

func ExistTagByName(name string) (*Tag, error) {
	var tag Tag
	err := db.Where("tag_name = ? ", name).First(&tag).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, err
	}
	return &tag, nil
}
