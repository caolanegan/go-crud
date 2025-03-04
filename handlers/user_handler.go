package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/caolanegan/go-crud/storage"
	"github.com/gocql/gocql"
)

// CreateUserHandler handles user creation
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newUser, err := storage.CreateUser(user.Name, user.Email)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// GetUserHandler retrieves a user by ID
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	parsedID, err := gocql.ParseUUID(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	user, err := storage.GetUser(parsedID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func ListAllUserHandlers(w http.ResponseWriter, r *http.Request) {
	users, err := storage.GetAllUsers()

	if err != nil {
		http.Error(w, "There was an error retrieving the list of users,", http.StatusNoContent)
		return
	}

	json.NewEncoder(w).Encode(users)
}
