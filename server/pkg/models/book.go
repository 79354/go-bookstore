package models

import (
	"go-bookstore/server/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct{
	gorm.Model // gives sturcture, helps store stuff in database

	Name 		string `json:"name"`
	Author 		string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

/*
	GORM - ORM object relational mapping
	Helps handle the sql queries
*/

func (b *Book) AddBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func BookByID(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db = db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func RemoveBook(Id int64) Book{
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}