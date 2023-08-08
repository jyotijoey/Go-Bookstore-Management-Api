package controllers

import (
	"encoding/json"
	"go_bookstore_management/pkg/models"
	"net/http"
)

var NewBooks models.Book

var (
	GetBook = func(w http.ResponseWriter, r *http.Request)  {
		newBooks := models.GetAllBooks()
		res, _ := json.Marshal(newBooks)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
	CreateBook = func(w http.ResponseWriter, r *http.Request)  {

	}
	GetBookById = func(w http.ResponseWriter, r *http.Request)  {

	}
	UpdateBookById = func(w http.ResponseWriter, r *http.Request)  {

	}
	DeleteBook = func(w http.ResponseWriter, r *http.Request)  {

	}
)