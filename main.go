package main

import (
	"log"
	"net/http"
	"os"

	"tripatra-test-go/db"
	"tripatra-test-go/graph"
	"tripatra-test-go/graph/generated"
	"tripatra-test-go/handlers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
)

func main() {
	db.Connect()

	r := mux.NewRouter()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	r.Use(handlers.CORSv2)

	r.HandleFunc("/login", handlers.Login).Methods("POST", "OPTIONS")
	r.Handle("/query", handlers.ValidateToken(srv))
	r.Handle("/playground", playground.Handler("GraphQL playground", "/query"))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 for local development
	}
	// Start server
	log.Printf("Server started on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
