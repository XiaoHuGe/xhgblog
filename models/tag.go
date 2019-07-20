package models

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model
	TagName    string //`json:"tag_name"`
	//CreatedBy  string `json:"created_by"`
	//ModifiedBy string `json:"modified_by"`
}

func GetTags(pageNum, pageSize int, maps map[string]interface{}) ([]Tag, error) {
	var (
		tags []Tag
		err error
	)
	if pageNum > 0 && pageSize > 0 {
		err = db.Where(maps).Find(&tags).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&tags).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tags, nil
}

func GetTag(id int) (Tag, error) {
	var tag Tag
	err := db.Where("id = ? ", id).Find(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tag, err
	}
	return tag, nil
}

func GetTagByName(name string) (Tag, error) {
	var tag Tag
	err := db.Where("tag_name = ? ", name).Find(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tag, err
	}
	return tag, nil
}

func DeleteTag(id int) (error) {
	err := db.Where("id = ?", id).Delete(&Tag{}).Error
	if err != nil {
		return err
	}
	return nil
}

func EditTag(id int, data interface{}) (error) {
	err := db.Model(&Tag{}).Where("id = ?", id).Update(data).Error
	if err != nil {
		return err
	}
	return nil
}

func AddTag(tag *Tag) (error) {
	err := db.Create(tag).Error
	if err != nil {
		return err
	}
	return nil
}

func GetTagTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Tag{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func ExistTagByID(id int) (bool, error) {
	err := db.Where("id = ? ", id).First(&Tag{}).Error
	if err != nil && err == gorm.ErrRecordNotFound { // 错误不为空且为未找到时返回false
		return false, err
	}
	return true, nil
}

func ExistTagByName(name string) (*Tag, error) {
	var tag Tag
	err := db.Where("tag_name = ? ", name).First(&tag).Error
	if err != nil && err == gorm.ErrRecordNotFound  {
		return nil, err
	}
	return &tag, nil
}