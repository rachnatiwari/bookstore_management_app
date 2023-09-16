package models

import (
	"log"

	"github.com/rachnatiwari/bookstore_management_app/pkg/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	result := db.Create(&b)
	if result.Error != nil {
		log.Fatal("error in creating record")
		panic("error creating data record")
	}
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) *Book {
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return &book
}

func UpdateBook(Id int64) *Book {
	var originalBook Book
	db.Where("ID=?", Id).Find(&originalBook)
	return &originalBook
}
