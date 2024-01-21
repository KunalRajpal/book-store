package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KunalRajpal/bookstore/pkg/models"
	"github.com/KunalRajpal/bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

//var NewBook models.Book //do you need this?

func GetBook(w http.ResponseWriter, r *http.Request) {
    newBooks := models.GetAllBooks()
    res, err := json.Marshal(newBooks)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    bookId := vars["bookId"]
    ID, err := strconv.ParseInt(bookId, 0, 0)
    if err != nil {
        http.Error(w, "Error while parsing book ID", http.StatusBadRequest)
        return
    }
    bookDetails, _ := models.GetBookById(ID)
    res, err := json.Marshal(bookDetails)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
    newBook := &models.Book{}
    err := utils.ParseBody(r, newBook)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    b := newBook.CreateBook()
    res, err := json.Marshal(b)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    bookId := vars["bookId"]
    ID, err := strconv.ParseInt(bookId, 0, 0)
    if err != nil {
        http.Error(w, "Error while parsing book ID", http.StatusBadRequest)
        return
    }

    book := models.DeleteBookById(ID)
    res, err := json.Marshal(book)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
    updateBook := &models.Book{}
    err := utils.ParseBody(r, updateBook)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    vars := mux.Vars(r)
    bookId := vars["bookId"]
    ID, err := strconv.ParseInt(bookId, 0, 0)
    if err != nil {
        http.Error(w, "Error while parsing book ID", http.StatusBadRequest)
        return
    }
    bookDetails, db := models.GetBookById(ID)
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
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}
