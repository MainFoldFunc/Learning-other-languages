package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/MainFoldFunc/Learning-other-languages/Go/Projects/Project_5_Book_store_managment/PKG/routes/bookstore-routes.go"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r) // Ensure this function is implemented in the routes package
	http.Handle("/", r)

	log.Println("Starting server on localhost:9010...")
	if err := http.ListenAndServe("localhost:9010", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

