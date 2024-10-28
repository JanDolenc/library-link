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
	ID         string `json:"id"`
	First_name string `json:"name"`
	Last_name  string `json:"surname"`
}

type usersResponse struct {
	Users []user `json:"users"`
}

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Quantity string `json:"quantity"`
}

type booksResponse struct {
	Books []book `json:"books"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a GET request at path /")
	w.Write([]byte("Welcome Library Link user. See /docs for more information."))
}

func queryForUsers() ([]user, error) {
	rows, _ := db.Query(context.Background(), "SELECT * FROM users")
	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[user])

	if err != nil {
		return nil, fmt.Errorf("Failed to query users: %w", err)
	}

	return users, nil
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a GET request at path /users")

	users, err := queryForUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response := usersResponse{
		Users: users,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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

func queryForBooks() ([]book, error) {
	rows, _ := db.Query(context.Background(), "SELECT * FROM books WHERE quantity != 0")
	books, err := pgx.CollectRows(rows, pgx.RowToStructByName[book])

	if err != nil {
		return nil, fmt.Errorf("Failed to query books: %w", err)
	}

	return books, nil
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a GET request at path /books")

	books, err := queryForBooks()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response := booksResponse{
		Books: books,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
