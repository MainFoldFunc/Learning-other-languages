package PKG

import "github.com/jinzhu/gorm"

type Book struct {
    gorm.Model
    Name        string `json:"name"`
    Author      string `json:"author"`
    Publication string `json:"publication"`
}

func init() {
    Connect()
    db := GetDB()
    db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
    db := GetDB()
    db.NewRecord(b)
    db.Create(&b)
    return b
}

func GetAllBooks() []Book {
    var books []Book
    db := GetDB()
    db.Find(&books)
    return books
}

func GetBookByID(ID int64) (*Book, *gorm.DB) {
    var getBook Book
    db := GetDB()
    db = db.Where("ID = ?", ID).Find(&getBook)
    return &getBook, db
}

func DeleteBook(ID int64) Book {
    var book Book
    db := GetDB()
    db.Where("ID = ?", ID).Delete(&book)
    return book
}

