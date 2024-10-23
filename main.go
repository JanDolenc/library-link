package main

import (
	"fmt"
	"net/http"
)


func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome Library Link user. See /docs for more information."))
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Starting server on port :8081")
	server.ListenAndServe()
}
