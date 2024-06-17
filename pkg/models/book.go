package models

import (
	"github.com/jinzhu/gorm"
	"go_bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:"json:""name"`
	Author   string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getbook Book
	db := db.Where("id = ?", Id).Find(&getbook)
	return &getbook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("id = ?", Id).Delete(&book)
	return book
}