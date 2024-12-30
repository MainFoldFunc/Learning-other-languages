package PKG

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

// GetBook handles GET requests to retrieve all books
func GetBook(w http.ResponseWriter, r *http.Request) {
    books := GetAllBooks()
    res, err := json.Marshal(books)
    if err != nil {
        http.Error(w, "Error marshalling data", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

// GetBookByID handles GET requests to retrieve a book by its ID
func GetBookBYID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    bookID := vars["bookID"]
    ID, err := strconv.ParseInt(bookID, 0, 64)
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }

    book, _ := GetBookByID(ID)
    if book.ID == 0 {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }

    res, err := json.Marshal(book)
    if err != nil {
        http.Error(w, "Error marshalling data", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

// CreateBook handles POST requests to create a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
    createBook := &Book{}
    ParseBody(r, createBook)

    newBook := createBook.CreateBook()
    res, err := json.Marshal(newBook)
    if err != nil {
        http.Error(w, "Error marshalling data", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

// UpdateBook handles PUT requests to update an existing book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
    var updateBook Book
    ParseBody(r, &updateBook)

    vars := mux.Vars(r)
    bookID := vars["bookID"]
    ID, err := strconv.ParseInt(bookID, 0, 64)
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }

    bookDetails, db := GetBookByID(ID)
    if bookDetails.ID == 0 {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }

    if updateBook.Name != "" {
        bookDetails.Name = updateBook.Name
    }
    if updateBook.Author != "" {
        bookDetails.Author = updateBook.Author
    }
    if updateBook.Publication != "" {
        bookDetails.Publication = updateBook.Publication
    }

    db.Save(&bookDetails)

    res, err := json.Marshal(bookDetails)
    if err != nil {
        http.Error(w, "Error marshalling data", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

// DeleteBook handles DELETE requests to remove a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    bookID := vars["bookID"]
    ID, err := strconv.ParseInt(bookID, 0, 64)
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }

    deletedBook := DeleteBook(ID)
    if deletedBook.ID == 0 {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }

    res, err := json.Marshal(deletedBook)
    if err != nil {
        http.Error(w, "Error marshalling data", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

