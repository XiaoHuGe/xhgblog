package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
	"xhgblog/utils/setting"
)

var db *gorm.DB

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createAt"`
	UpdatedAt time.Time  `json:"updateAt"`
	DeletedAt *time.Time `json:"deleteAt"`
}

func Setup() {
	var err error
	db, err = gorm.Open(setting.AppSetting.Database.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.AppSetting.Database.User,
		setting.AppSetting.Database.Password,
		setting.AppSetting.Database.Host,
		setting.AppSetting.Database.Name))
	//defer db.Close()

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.AppSetting.Database.TablePrefix + defaultTableName
	}

	db.LogMode(true)
	db.SingularTable(true)

	db.DB().SetConnMaxLifetime(time.Minute * 3)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&Article{})
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Page{})
	db.AutoMigrate(&Category{})
}
