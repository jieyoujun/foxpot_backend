package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver
)

// DB ...
var DB *gorm.DB

// InitDB ...
func InitDB() (err error) {
	DB, err = gorm.Open("mysql", "root:1212@tcp(localhost:3306)/gorm_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil
	}
	return nil
}

// InitDBAndMigrateAll ...
func InitDBAndMigrateAll() {
	if err := InitDB(); err != nil {
		log.Fatalln("Failed to connect to mysql:", err)
	}
	DB.AutoMigrate(&User{})
}
