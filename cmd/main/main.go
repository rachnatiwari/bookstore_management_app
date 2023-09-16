package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rachnatiwari/bookstore_management_app/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.BooksRoute(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
