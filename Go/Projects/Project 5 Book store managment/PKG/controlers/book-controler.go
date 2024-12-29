package controlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/MainFoldFunc/Learning-other-languages/go/Projects/Project 5 Bookstore managment/PKG/utils"
	"github.com/MainFoldFunc/Learning-other-languages/go/Projects/Project 5 Bookstore managment/PKG/models"
)

var NewBook models.Book

func GetBook(w *http.ResponseWriter, r *http.Request) {
	newbooks := models.GetAllBooks()
	res, _ := json.Marshal(newbooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w *http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["bookID"]
	ID, err :=strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Invalid Request")
	}
	book, _ := models.GetBookByID(int64(ID))
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w *http.ResponseWriter, r *http.Request) {
	Createbook := &models.Book{}
	utils.Parsebody(r, Createbook)
	b := Createbook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w *http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Invalid Request")
	}
	book := models.DeleteBook(int64(ID))
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w *http.ResponseWriter, r *http.Request) {
	var UpdateBook models.Book
	utils.ParseInt(r, UpdateBook)
	vars := mux.Vars(r)
	bookID := vars["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	bookdetails, db := models.GetBookByID(int64(ID))
	if updateBook.Name != "" {
		bookdetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookdetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" Invalid Request{
		bookdetails.Publication = updateBook.Publication
	}
	db.Save(&bookdetails)
	res, _ := json.Marshal(bookdetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


