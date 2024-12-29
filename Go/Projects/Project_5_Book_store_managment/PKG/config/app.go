package config

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
    db *gorm.DB
)

func Connect() {
    d, err := gorm.Open("mysql", "MainFoldFunc:2010/simplerest?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic("Failed to connect to database!")
    }
    db = d
}

func GetDB() *gorm.DB {
    return db
}