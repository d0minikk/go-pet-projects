package main

import (
	"fmt"
	"net/http"

	"ex3-rest/handlers"
	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", handlers.ReadBooksHandler).Methods("GET")
	router.HandleFunc("/books", handlers.CreateBookHandler).Methods("POST")
	router.HandleFunc("/books/{id:[0-9]+}", handlers.ReadBookHandler).Methods("GET")
	router.HandleFunc("/books/{id:[0-9]+}", handlers.UpdateBookHandler).Methods("PUT")

	port := "8080"
	fmt.Printf("Server is running on port %s\n", port)
	http.ListenAndServe(":"+port, router)
}
