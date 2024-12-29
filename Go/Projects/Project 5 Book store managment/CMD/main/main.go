package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"error"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/MainFoldFunc/Learning-other-languages/go/Projects/Project 5 Bookstore managment/PKG/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
