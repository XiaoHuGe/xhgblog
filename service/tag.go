package service

import (
	"fmt"
	"xhgblog/models"
)

type GetTagService struct {
	TagName string
	//State    int
	PageNum  int
	PageSize int
}

func (this *GetTagService) GetAll() ([]models.Tag, error) {
	maps := make(map[string]interface{})
	if this.TagName != "" {
		maps["name"] = this.TagName
	}
	tags, err := models.GetTags(this.PageNum, this.PageSize, maps)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (this *GetTagService) Count() (int, error) {
	return models.GetTagTotal(make(map[string]interface{}))
}

type AddTagService struct {
	TagName string `form:"value" json:"value" binding:"required,min=2,max=10"`
	//CreatedBy string `form:"created_by" json:"created_by" binding:"required,min=2,max=10"`
}

func (this *AddTagService) AddTag() (*models.Tag, error) {

	tag, err := models.ExistTagByName(this.TagName)
	if err == nil {
		fmt.Println("存在此标签")
		tag.Total += 1
		models.EditTag(int(tag.ID), tag)
		return tag, nil
	}

	tag = &models.Tag{
		TagName: this.TagName,
		Total:   1,
		//CreatedBy: this.CreatedBy,
	}
	err = models.AddTag(tag)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

type EditTagService struct {
	TagName string `form:"tag_name" json:"tag_name" binding:"required,min=2,max=10"`
	Total   int
	//ModifiedBy string `form:"modified_by" json:"modified_by" binding:"required,min=2,max=10"`
}

func (this *EditTagService) EdidTag(id int) error {
	tag := &models.Tag{
		TagName: this.TagName,
		Total:   this.Total,
		//ModifiedBy: this.ModifiedBy,
	}
	return models.EditTag(id, tag)
}

func CheckTagByID(id int) (*models.Tag, error) {
	return models.ExistTagByID(id)
}

func DeleteTag(id int) error {
	return models.DeleteTag(id)
}
