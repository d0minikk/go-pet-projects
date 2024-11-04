package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	
	"ex3-rest/models"
	"ex3-rest/utils"
	
	"github.com/gorilla/mux"
)

var books = []models.Book{
  {ID: "1", Title: "Book One", Author: "Author A", Price: 19.99},
  {ID: "2", Title: "Book Two", Author: "Author B", Price: 29.99},
}

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}
	defer r.Body.Close()

	if book.Title == "" || book.Author == "" {
		utils.WriteJSONError(w, http.StatusBadRequest, "Missing required fields: title or author")
		return
	}

	book.ID = fmt.Sprintf("%d", len(books)+1)
	books = append(books, book)

	fmt.Printf("Received new book: %+v\n", book)

	utils.WriteJSONResponse(w, http.StatusCreated, book)
}

func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updatedBook models.Book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}
	defer r.Body.Close()

	if updatedBook.Title == "" || updatedBook.Author == "" {
		utils.WriteJSONError(w, http.StatusBadRequest, "Missing required fields: title or author")
		return
	}

	for i, book := range books {
		if book.ID == id {
			updatedBook.ID = id
			books[i] = updatedBook
			fmt.Printf("Updating book with ID %s: %+v\n", id, updatedBook)
			utils.WriteJSONResponse(w, http.StatusOK, updatedBook)
			return
		}
	}

	utils.WriteJSONError(w, http.StatusNotFound, "Book not found")
}

func ReadBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, book := range books {
		if book.ID == id {
			utils.WriteJSONResponse(w, http.StatusOK, book)
			return
		}
	}

	utils.WriteJSONError(w, http.StatusNotFound, "Book not found")
}

func ReadBooksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Starting ReadBooksHandler")
	utils.WriteJSONResponse(w, http.StatusOK, books)
}
