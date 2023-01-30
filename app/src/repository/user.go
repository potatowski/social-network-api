package repository

import (
	"database/sql"
	"social-api/src/model"
)

// User represents an user repository
type User struct {
	db *sql.DB
}

// NewRepositoryUser create an user repository
func NewRepositoryUser(db *sql.DB) *User {
	return &User{db}
}

// Create insert an user in database
func (userRepository User) Create(user model.User) (uint64, error) {
	query := "INSERT INTO user (name, username, email, password) VALUES (?, ?, ?, ?)"
	statement, err := userRepository.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(userID), nil
}
