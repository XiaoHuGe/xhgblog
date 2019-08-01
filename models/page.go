package models

type Page struct {
	Model
	Title       string `json:"title"`
	Content     string `json:"content"`
	IsPublished bool   `json:"is_published"`
}

func AddPage(page *Page) error {
	return db.Create(page).Error
}

func GetPages() ([]*Page, error) {
	var page []*Page
	err := db.Find(&page).Error
	return page, err
}


func EditPage(id int, p *Page) error {
	var page Page
	err := db.Where("id = ? ", id).First(&page).Update(p).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPageByID(id int) (*Page, error) {
	var page Page
	err := db.Where("id = ?", id).First(&page).Error
	return &page, err
}

func GetPageByTitle(title string) (*Page, error) {
	var page Page
	err := db.Where("title = ?", title).First(&page).Error
	return &page, err
}

func DeletePageByID(id int) error {
	return db.Where("id = ?", id).Delete(&Page{}).Error
}

func ExistPageByID(id int) (bool, error) {
	var page Page
	err := db.Where("id = ? ", id).First(&page).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
