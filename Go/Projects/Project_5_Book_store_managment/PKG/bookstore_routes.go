package PKG

import (
    "github.com/gorilla/mux"
)

// RegisterBookStoreRoutes defines routes for the book store
var RegisterBookStoreRoutes = func(router *mux.Router) {
    router.HandleFunc("/books", GetBook).Methods("GET")
    router.HandleFunc("/books/{bookID}", GetBookByID).Methods("GET")
    router.HandleFunc("/books", CreateBook).Methods("POST")
    router.HandleFunc("/books/{bookID}", UpdateBook).Methods("PUT")
    router.HandleFunc("/books/{bookID}", DeleteBook).Methods("DELETE")
}

