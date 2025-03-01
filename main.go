package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/caolanegan/go-crud/storage" // Import database package
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Connect to Cassandra
	storage.ConnectDatabase()
	defer storage.Session.Close() // Ensure session closes when app stops

	r := chi.NewRouter()
	r.Use(middleware.Logger) // Logs incoming requests

	// Health check route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Go CRUD API!"))
	})

	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
