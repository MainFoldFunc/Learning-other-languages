package PKG

import (
	"github.com/gorilla/mux"
)
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controlers.GetBook).Methods("GET")
	router.HandleFunc("/books/{bookID}", controlers.GetBookByID).Methods("GET")
	router.HandleFunc("/books", controlers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{bookID}", controlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookID}", controlers.DeleteBook).Methods("DELETE")
}

