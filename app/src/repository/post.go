package repository

import (
	"database/sql"
	"social-api/src/model"
)

// Post represents a post repository
type Post struct {
	db *sql.DB
}

// NewRepositoryPost creates a new post repository
func NewRepositoryPost(db *sql.DB) *Post {
	return &Post{db}
}

// Create creates a new post in database
func (repository Post) Create(post model.Post) (uint64, error) {
	statement, err := repository.db.Prepare("INSERT INTO post (uuid, title, body, user_id) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.UUID, post.Title, post.Body, post.UserID)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}
