package models

import "github.com/jinzhu/gorm"

type Category struct {
	Model
	CategoryName string //`json:"category_name"`
	Total        int    `json:"total" gorm:"-"` // article num
}

func AddCategory(name string) (*Category, error) {
	var c Category
	err := db.Where(Category{CategoryName: name}).FirstOrCreate(&c).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &c, nil
}

func GetCategoryById(id string) (*Category, error) {
	var c Category
	err := db.Where("id = ?)", id).Find(&c).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &c, err
	}
	return &c, nil
}

func GetCategorys() ([]*Category, error) {
	var (
		categorys []*Category
		err       error
	)
	rows, err := db.Raw("select c.*,count(*) total from xhgblog_category c join xhgblog_article p on p.category_id = c.id group by p.category_id").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var category Category
		db.ScanRows(rows, &category)
		categorys = append(categorys, &category)
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return categorys, nil
}
