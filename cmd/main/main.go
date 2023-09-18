package main

import (
	"log"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/rachnatiwari/bookstore_management_app/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.BooksRoute(r)
	r.Handle("/swagger_main.yml", http.FileServer(http.Dir("./")))
	opts := middleware.SwaggerUIOpts{SpecURL: "swagger_main.yml"}
	sh := middleware.SwaggerUI(opts, nil)
	r.Handle("/docs", sh)
	r.Handle("/", http.RedirectHandler("/docs", http.StatusSeeOther))
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
