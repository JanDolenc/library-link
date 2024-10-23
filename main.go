package main

import (
	"fmt"
	"log"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a request at path /")
	w.Write([]byte("Welcome Library Link user. See /docs for more information."))
}


func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", greet)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Starting server on port :8081")
	server.ListenAndServe()
}
