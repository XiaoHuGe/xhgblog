package service

import "xhgblog/models"

type PageService struct {
	Title       string `form:"title"`
	Content     string `form:"content"`
	IsPublished string `form:"is_published"`
}

func (this *PageService) AddPage() error {
	published := this.IsPublished == "on"
	page := &models.Page{
		Title:       this.Title,
		Content:     this.Content,
		IsPublished: published,
	}
	return models.AddPage(page)
}

func (this *PageService) EditPage(id int) error {
	published := this.IsPublished == "on"
	page := &models.Page{
		Title:       this.Title,
		Content:     this.Content,
		IsPublished: published,
	}
	return models.EditPage(id, page)
}

func (this *PageService) GetPageByID(id int) (*models.Page, error)  {
	return models.GetPageByID(id)
}

func (this *PageService) GetPageByTitle(title string) (*models.Page, error)  {
	return models.GetPageByTitle(title)
}

func (this *PageService) DeletePageByID(id int) (error)  {
	return models.DeletePageByID(id)
}

func (this *PageService)CheckPageByID(id int) (bool, error) {
	return models.ExistPageByID(id)
}
