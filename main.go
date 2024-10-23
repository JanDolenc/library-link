package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type user struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

var users = []user{
	{ID: "1", Name: "Jan", Surname: "Dolenc"},
	{ID: "2", Name: "Manca", Surname: "Cater"},
	{ID: "3", Name: "Ajda", Surname: "Maček"},
	{ID: "4", Name: "Marko", Surname: "Pisk"},
	{ID: "5", Name: "Rok", Surname: "Puntar"},
	{ID: "6", Name: "Boštjan", Surname: "Dolenec"},
	{ID: "7", Name: "Janez", Surname: "Možina"},
}

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Quantity string `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Don Quixote", Quantity: "15"},
	{ID: "2", Title: "A Tale of Two Cities", Quantity: "20"},
	{ID: "3", Title: "War and Peace", Quantity: "12"},
	{ID: "4", Title: "Moby-Dick", Quantity: "18"},
	{ID: "5", Title: "The Count of Monte Cristo", Quantity: "22"},
	{ID: "6", Title: "Jane Eyre", Quantity: "15"},
	{ID: "7", Title: "Wuthering Heights", Quantity: "10"},
	{ID: "8", Title: "Great Expectations", Quantity: "25"},
}

func greet(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a request at path /")
	w.Write([]byte("Welcome Library Link user. See /docs for more information."))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a request at path /users")

	// convert slice to json format
	usersJSON, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set header Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON repsonse to the client
	w.Write(usersJSON)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a request at path /books")

	booksJSON, err := json.Marshal(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(booksJSON)
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", greet)
	router.HandleFunc("GET /users", getUsers)
	router.HandleFunc("GET /books", getBooks)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Starting server on port :8081")
	server.ListenAndServe()
}
