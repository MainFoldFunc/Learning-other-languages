package routes

import (
	"github.com/gorilla/mux"
	"github.com/MainFoldFunc/Learning-other-languages/Go/Projects/Project5_Book_store_Management/PKG"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controlers.GetBook).Methods("GET")
	router.HandleFunc("/books/{bookID}", controlers.GetBookByID).Methods("GET")
	router.HandleFunc("/books", controlers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{bookID}", controlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookID}", controlers.DeleteBook).Methods("DELETE")
}

