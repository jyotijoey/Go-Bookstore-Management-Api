package controllers

import (
	"encoding/json"
	"go_bookstore_management/pkg/models"
	"net/http"
)

var NewBooks models.Book

// var (
	func GetBook(w http.ResponseWriter, r *http.Request)  {
		newBooks := models.GetAllBooks()
		res, _ := json.Marshal(newBooks)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
	func CreateBook(w http.ResponseWriter, r *http.Request)  {

	}
	func GetBookById(w http.ResponseWriter, r *http.Request)  {

	}
	func UpdateBookById(w http.ResponseWriter, r *http.Request)  {

	}
	func DeleteBook(w http.ResponseWriter, r *http.Request)  {

	}
// )