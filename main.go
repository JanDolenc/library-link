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

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", greet)
	router.HandleFunc("GET /users", getUsers)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Starting server on port :8081")
	server.ListenAndServe()
}
