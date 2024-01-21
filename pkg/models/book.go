package models

import (
	"log"

	"github.com/KunalRajpal/bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct{
	gorm.Model
    Name        string `json:"name"`
    Author      string `json:"author"`
    Publication string `json:"publication"`
}

func init(){
	err :=config.Connect()
	if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db:=db.Where("ID=?",Id).Find(&getBook)
	return &getBook, db
}

func DeleteBookById(ID int64) Book{
	var book Book
	db.Where("ID=?", ID).Delete(book) //&book?
	return book
}