package controlers

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
    "strconv"
    "github.com/MainFoldFunc/Learning-other-languages/go/Projects/Project_5_Bookstore_managment/PKG/utils"
    "github.com/MainFoldFunc/Learning-other-languages/go/Projects/Project_5_Bookstore_managment/PKG/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
    newbooks := models.GetAllBooks()
    res, _ := json.Marshal(newbooks)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    bookID := vars["bookID"]
    ID, err := strconv.ParseInt(bookID, 0, 0)
    if err != nil {
        fmt.Println("Invalid Request")
    }
    book, _ := models.GetBookByID(int64(ID))
    res, _ := json.Marshal(book)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
    Createbook := &models.Book{}
    utils.ParseBody(r, Createbook)
    b := Createbook.CreateBook()
    res, _ := json.Marshal(b)
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    bookID := vars["bookID"]
    ID, err := strconv.ParseInt(bookID, 0, 0)
    if err != nil {
        fmt.Println("Invalid Request")
    }
    book := models.DeleteBook(int64(ID))
    res, _ := json.Marshal(book)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
    var UpdateBook models.Book
    utils.ParseBody(r, &UpdateBook)
    vars := mux.Vars(r)
    bookID := vars["bookID"]
    ID, err := strconv.ParseInt(bookID, 0, 0)
    if err != nil {
        fmt.Println("error while parsing")
    }
    bookdetails, db := models.GetBookByID(int64(ID))
    if UpdateBook.Name != "" {
        bookdetails.Name = UpdateBook.Name
    }
    if UpdateBook.Author != "" {
        bookdetails.Author = UpdateBook.Author
    }
    if UpdateBook.Publication != "" {
        bookdetails.Publication = UpdateBook.Publication
    }
    db.Save(&bookdetails)
    res, _ := json.Marshal(bookdetails)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}
