package models

import "github.com/gocql/gocql"

// User represents a user in the system
type User struct {
	ID    gocql.UUID `json:"id"`
	Name  string     `json:"name"`
	Email string     `json:"email"`
}
