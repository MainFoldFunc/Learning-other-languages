package PKG

import (
    "github.com/gorilla/mux"
)

// RegisterBookStoreRoutes defines routes for the book store
var RegisterBookStoreRoutes = func(router *mux.Router) {
    router.HandleFunc("/books", GetBooK).Methods("GET")
    router.HandleFunc("/books/{bookID}", GetBookBYID).Methods("GET")
    router.HandleFunc("/books", CreateBooK).Methods("POST")
    router.HandleFunc("/books/{bookID}", UpdateBooK).Methods("PUT")
    router.HandleFunc("/books/{bookID}", DeleteBooK).Methods("DELETE")
}

