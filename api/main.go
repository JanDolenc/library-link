package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

type newUser struct {
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
	w.Write([]byte("Welcome Library Link user."))
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

func insertUser(user newUser) error {
	_, err := db.Exec(context.Background(), "INSERT INTO users (first_name, last_name) VALUES($1, $2)", user.First_name, user.Last_name)

	if err != nil {
		return fmt.Errorf("Unable to insert row: %w", err)
	}

	return nil
}

func createUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a POST request at path /users")

	var user newUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if user.First_name == "" {
		http.Error(w, "Error: Property 'name' is required to create a new user and must have a non empty value.", http.StatusBadRequest)
		return
	}
	if user.Last_name == "" {
		http.Error(w, "Error: Property 'surname' is required to create a new user and must have a non empty value.", http.StatusBadRequest)
		return
	}

	log.Printf("Creating new user: Name %s, Surname %s\n", user.First_name, user.Last_name)

	if err := insertUser(user); err != nil {
		log.Printf("Failed inserting new user. %v\n", err)
		http.Error(w, "Error: Could not insert the new user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("User %s %s created successfully", user.First_name, user.Last_name),
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
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("HOST_DB")
	port := os.Getenv("PORT_DB")
	database := os.Getenv("POSTGRES_DB")

	database_url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)

	var err error
	// connect to db with pgx - postgres driver
	db, err = pgx.Connect(context.Background(), database_url)
	if err != nil {
		return fmt.Errorf("Unable to connect to database %w", err)
	}
	return nil
}

func main() {
	// load environment variables from file - for running on localhost
	if err := godotenv.Load("../.env"); err != nil {
		log.Printf("Important for LOCALHOST. Failed to load .env file. %v If running in Docker you can ignore this.\n\n", err)
	}

	if err := initDB(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize db: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	router := http.NewServeMux()

	router.HandleFunc("GET /", greet)
	router.HandleFunc("GET /users", getUsers)
	router.HandleFunc("POST /users", createUser)
	router.HandleFunc("GET /books", getBooks)

	apiPort := os.Getenv("PORT_API")
	apiPortHost := os.Getenv("HOST_PORT_API")
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", apiPort),
		Handler: router,
	}

	fmt.Printf("Starting server on port :%s->%s\n", apiPortHost, apiPort)
	server.ListenAndServe()
}
