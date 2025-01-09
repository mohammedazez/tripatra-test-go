package main

import (
	"log"
	"net/http"

	"tripatra-test-go/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// Start server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
