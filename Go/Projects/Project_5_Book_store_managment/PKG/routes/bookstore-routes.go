package routes

import (
	"github.com/gorilla/mux"
	"github.com/MainFoldFunc/Learning-other-languages/go/Projects/Project_5_Bookstore_managment/PKG/controlers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controlers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controlers.GetBookID).Methods("GET")
	router.HandleFunc("/books", controlers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", controlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controlers.DeleteBook).Methods("DELETE")
}
