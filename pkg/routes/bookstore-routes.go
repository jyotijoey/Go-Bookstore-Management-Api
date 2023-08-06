package routes

import (
	"go_bookstore_management/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookstoreRoute = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBookById).Methods("PUT")
	router.HandleFunc("/book/{bookID}", controllers.DeleteBook).Methods("DELETE")
}