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

	rows, err := userRepository.db.Query("SELECT id, name, username, email, created FROM user WHERE name LIKE ? or username LIKE ? AND removed <> 1", search, search)
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
	row, err := userRepository.db.Query("SELECT id, name, username, email, created FROM user WHERE id = ? AND removed <> 1", userID)
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

// Delete remove an user in database
func (userRepository User) Delete(userID uint64) error {
	statement, err := userRepository.db.Prepare("UPDATE user SET removed = 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID); err != nil {
		return err
	}

	return nil
}

// SearchByEmail try find an user in database by email
func (userRepository User) SearchByEmail(email string) (model.User, error) {
	row, err := userRepository.db.Query("SELECT id, password FROM user WHERE email = ? AND removed <> 1", email)
	var user model.User
	if err != nil {
		return user, err
	}
	defer row.Close()

	if row.Next() {
		if err = row.Scan(
			&user.ID,
			&user.Password,
		); err != nil {
			return user, err
		}
	}

	return user, nil
}

// Follow insert a new follow in database
func (userRepository User) Follow(userID, followerID uint64) error {
	statement, err := userRepository.db.Prepare("INSERT IGNORE INTO follower (user_id, follower_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// Unfollow remove a follow in database
func (userRepository User) Unfollow(userID, followerID uint64) error {
	statement, err := userRepository.db.Prepare("DELETE FROM follower WHERE user_id = ? AND follower_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// SearchFollowers try find followers of an user in database
func (userRepository User) SearchFollowers(userId uint64) ([]model.User, error) {
	rows, err := userRepository.db.Query(`
		SELECT u.id, u.name, u.username, u.email, u.created
		FROM user u
		INNER JOIN follower f ON u.id = f.follower_id
		WHERE f.user_id = ?
		AND u.removed <> 1
	`, userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []model.User
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

		followers = append(followers, user)
	}

	return followers, nil
}
