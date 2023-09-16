package routes

import (
	"github.com/gorilla/mux"
	controller "github.com/rachnatiwari/bookstore_management_app/pkg/controllers"
)

var BooksRoute = func(router *mux.Router) {
	router.HandleFunc("/books", controller.ShowAllBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controller.ShowBookDetail).Methods("GET")
	router.HandleFunc("/book", controller.AddNewBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controller.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{bookId}", controller.UpdateBook).Methods("PUT")
}
