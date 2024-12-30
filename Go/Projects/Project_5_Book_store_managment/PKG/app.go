package PKG

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("sqlite3", "bookstore.db")
	if err != nil {
		panic("Failed to connect to database!")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

