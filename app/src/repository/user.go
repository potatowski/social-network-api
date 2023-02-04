package repository

import (
	"database/sql"
	"fmt"
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

// Search try find an user in database by username or name
func (userRepository User) Search(search string) ([]model.User, error) {
	search = fmt.Sprintf("%%%s%%", search)

	rows, err := userRepository.db.Query("SELECT id, name, username, email, created FROM user WHERE name LIKE ? or username LIKE ?", search, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.Created,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func (userRepository User) SearchById(userID uint64) (model.User, error) {
	row, err := userRepository.db.Query("SELECT id, name, username, email, created FROM user WHERE id = ?", userID)
	var user model.User
	if err != nil {
		return user, err
	}
	defer row.Close()

	if row.Next() {
		if err = row.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.Created,
		); err != nil {
			return user, err
		}
	}

	return user, nil
}

// Update change info of user in database
func (userRepository User) Update(userID uint64, user model.User) error {
	statement, err := userRepository.db.Prepare("UPDATE user SET name = ?, username = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Username, user.Email, userID); err != nil {
		return err
	}

	return nil
}
