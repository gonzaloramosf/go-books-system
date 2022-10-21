package main

import (
	"github.com/gorilla/mux"
	"go-book-managment/db"
	"go-book-managment/models"
	"go-book-managment/router"
	"log"
	"net/http"
)

func main() {
	db.ConnectionDB()

	db.DB.AutoMigrate(models.Book{})

	r := mux.NewRouter()
	router.BooksRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
