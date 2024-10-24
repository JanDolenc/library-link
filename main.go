package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
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

func createUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a POST request at path /users \n")

	queryParams := r.URL.RawQuery

	// log.Println("raw query values:", queryParams)

	queryParamsMap, err := url.ParseQuery(queryParams)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if queryParamsMap["name"] == nil || queryParamsMap["name"][0] == "" {
		http.Error(w, "Error: Mising query parameter. Parameter 'name' is required and must have a non empty value.", http.StatusInternalServerError)
		return
	}

	if queryParamsMap["surname"] == nil || queryParamsMap["surname"][0] == "" {
		http.Error(w, "Error: Mising query parameter. Parameter 'surname' is required and must have a non empty value.", http.StatusInternalServerError)
		return
	}

	log.Println("queryParamsMap", queryParamsMap)
	log.Println("accesing map value", queryParamsMap["name"], len(queryParamsMap["name"]))
	log.Printf("type of queryParamsMap[name] %T, %T, %v", queryParamsMap["name"], queryParamsMap["name"][0], len(queryParamsMap["name"][0]))

	log.Println("Hello", queryParamsMap["name"][0], queryParamsMap["surname"][0], "! You are now a registered user of Library Link.")
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
	router.HandleFunc("POST /users", createUser)
	router.HandleFunc("GET /books", getBooks)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Starting server on port :8081")
	server.ListenAndServe()
}
