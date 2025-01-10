package main

import (
	"log"
	"net/http"

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

	r.Use(handlers.CORS)

	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.Handle("/query", handlers.ValidateToken(srv))
	r.Handle("/playground", playground.Handler("GraphQL playground", "/query"))

	// Start server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
