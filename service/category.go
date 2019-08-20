package service

import "xhgblog/models"

type AddCategoryService struct {
	CategoryName string `form:"value" json:"value" binding:"required,min=2,max=10"`
	//CreatedBy string `form:"created_by" json:"created_by" binding:"required,min=2,max=10"`
}

func (this *AddCategoryService) AddCategory() (*models.Category, error) {

	tag, err := models.AddCategory(this.CategoryName)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func (this *AddCategoryService) GetCategorys() (*models.Category, error) {

	tag, err := models.AddCategory(this.CategoryName)
	if err != nil {
		return nil, err
	}
	return tag, nil
}
