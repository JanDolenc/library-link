package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var db *pgx.Conn

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
	log.Println("Received a GET request at path /")
	w.Write([]byte("Welcome Library Link user. See /docs for more information."))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a GET request at path /users")

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
	queryParamsMap, err := url.ParseQuery(queryParams)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if queryParamsMap["name"] == nil || queryParamsMap["name"][0] == "" {
		http.Error(w, "Error: Mising query parameter. Parameter 'name' is required and must have a non empty value.", http.StatusBadRequest)
		return
	}

	if queryParamsMap["surname"] == nil || queryParamsMap["surname"][0] == "" {
		http.Error(w, "Error: Mising query parameter. Parameter 'surname' is required and must have a non empty value.", http.StatusBadRequest)
		return
	}

	log.Printf("Creating new user: Name %s, Surname %s\n", queryParamsMap["name"][0], queryParamsMap["surname"][0])

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("User %s %s created successfully", queryParamsMap["name"][0], queryParamsMap["surname"][0]),
	})
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a GET request at path /books")

	booksJSON, err := json.Marshal(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(booksJSON)
}

func initDB() error {
	var err error
	// connect to db with pgx - postgres driver
	db, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return fmt.Errorf("Unable to connect to database %w", err)
	}
	return nil
}

func main() {
	// load environment variables from file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// log.Printf("Database url: %v\n", os.Getenv("DATABASE_URL"))

	if err := initDB(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize db: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	// Test query
	var greeting string
	err = db.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

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
