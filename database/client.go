package database

import (
	"github.com/Dwata-Tech/golang-test-task/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func Migrate() {
	Instance.AutoMigrate(&model.Article{})
	Instance.AutoMigrate(&model.Comment{})
	log.Println("Database Migration Completed...")
}

func CreateArticle(article *model.Article) *gorm.DB {
	return Instance.Create(&article)
}
