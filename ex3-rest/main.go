package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func ReadBookHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	fmt.Println("Requested book with id %s", id)
}

func ReadBooksHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Starting ReadBooksHandler")
}

func CreateBookHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Starting CreateBookHandler")
}

func UpdateBookHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Starting UpdateBookHandler")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books/{id:[0-9]+}", ReadBookHandler).Methods("GET")
	router.HandleFunc("/books/{id:[0-9]+}", UpdateBookHandler).Methods("PUT")
	router.HandleFunc("/books", ReadBooksHandler).Methods("GET")
	router.HandleFunc("/books", CreateBookHandler).Methods("POST")

	http.ListenAndServe(":80", router)
}
