package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-book-managment/db"
	"go-book-managment/models"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	json.NewDecoder(r.Body).Decode(&book)

	createdBook := db.DB.Create(&book)
	err := createdBook.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	db.DB.Find(&books)

	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	params := mux.Vars(r)
	db.DB.First(&book, params["id"])

	if book.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book not found"))
		return
	}

	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	params := mux.Vars(r)

	db.DB.First(&book, params["id"])

	if book.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book you are trying to update not found"))
		return
	}

	json.NewDecoder(r.Body).Decode(&book)

	db.DB.Unscoped().Save(book)
	w.WriteHeader(http.StatusOK)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	params := mux.Vars(r)

	db.DB.First(&book, params["id"])

	if book.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book you are trying to delete not found"))
		return
	}

	db.DB.Unscoped().Delete(book)
	w.WriteHeader(http.StatusOK)
}
