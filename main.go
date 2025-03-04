package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/caolanegan/go-crud/handlers"
	"github.com/caolanegan/go-crud/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Connect to Cassandra
	storage.ConnectDatabase()
	defer storage.Session.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Health check
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Go CRUD API!"))
	})

	// User Routes
	r.Post("/users", handlers.CreateUserHandler)      // Create user
	r.Get("/users", handlers.GetUserHandler)          // Get user by ID
	r.Get("/listUsers", handlers.ListAllUserHandlers) // List all users

	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
