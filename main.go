package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Volume string  `json:"volume"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"first name"`
	LastName  string `json:"last name"`
	Age       int    `json:"age"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func deleteBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			json.NewDecoder(r.Body).Decode(&book)
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	books = append(books, Book{ID: "1", Title: "Ramatohara", Volume: "Medium", Author: &Author{FirstName: "Ysmaeke", LastName: "Worku", Age: 45}})
	books = append(books, Book{ID: "2", Title: "Fkir Eskemekabr", Volume: "High", Author: &Author{FirstName: "Hadis", LastName: "Alemayew", Age: 60}})
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBookByID).Methods("GET")
	r.HandleFunc("/books", createBooks).Methods("POST")
	r.HandleFunc("/books/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBooks).Methods("DELETE")

	fmt.Println("server starting...")
	log.Fatal(http.ListenAndServe(":6000", r))
}
