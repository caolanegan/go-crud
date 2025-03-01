package storage

import (
	"github.com/caolanegan/go-crud/models"
	"github.com/gocql/gocql"
)

// CreateUser inserts a new user into Cassandra
func CreateUser(name, email string) (models.User, error) {
	id := gocql.TimeUUID()
	query := "INSERT INTO users (id, name, email) VALUES (?, ?, ?)"
	err := Session.Query(query, id, name, email).Exec()
	if err != nil {
		return models.User{}, err
	}

	return models.User{ID: id, Name: name, Email: email}, nil
}

// GetUser retrieves a user by ID
func GetUser(id gocql.UUID) (models.User, error) {
	var user models.User
	query := "SELECT id, name, email FROM users WHERE id = ? LIMIT 1"
	err := Session.Query(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// GetAllUsers retrieves all users
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	query := "SELECT id, name, email FROM users"

	iter := Session.Query(query).Iter()
	var id gocql.UUID
	var name, email string

	for iter.Scan(&id, &name, &email) {
		users = append(users, models.User{ID: id, Name: name, Email: email})
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return users, nil
}

// DeleteUser deletes a user by ID
func DeleteUser(id gocql.UUID) error {
	query := "DELETE FROM users WHERE id = ?"
	return Session.Query(query, id).Exec()
}
