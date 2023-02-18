package repository

import (
	"database/sql"
	"social-api/src/model"
)

// Post represents an post repository
type Post struct {
	db *sql.DB
}

// NewRepositoryPost creates a new post repository
func NewRepositoryPost(db *sql.DB) *Post {
	return &Post{db}
}

// Create creates a new post in database
func (postRepository Post) Create(post model.Post) (uint64, error) {
	statement, err := postRepository.db.Prepare("INSERT INTO post (uuid, title, body, user_id) values (?, ?, ?, ?)")
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

func (postRepository Post) SearchByUuid(uuid string) (model.Post, error) {
	var post model.Post
	rows, err := postRepository.db.Query(`
		SELECT 
			p.id, p.uuid, p.title, p.body, p.likes, p.created, p.user_id,
			u.id, u.name, u.username, u.created
		FROM post p INNER JOIN user u ON u.id = p.user_id
		WHERE p.uuid = ? AND p.removed <> 1
	`, uuid)
	if err != nil {
		return post, err
	}
	defer rows.Close()

	if rows.Next() {
		var user model.User
		if err = rows.Scan(
			&post.ID,
			&post.UUID,
			&post.Title,
			&post.Body,
			&post.Likes,
			&post.Created,
			&post.UserID,
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Created,
		); err != nil {
			return post, err
		}
		post.User = &user
	}

	return post, nil
}

func (postRepository Post) SearchByUser(userId uint64) ([]model.Post, error) {
	rows, err := postRepository.db.Query(`
		SELECT DISTINCT 
			p.uuid, p.title, p.body, p.likes, p.created,
			u.id, u.name, u.username
		FROM post p 
		INNER JOIN user u ON u.id = p.user_id
		INNER JOIN follower f ON p.user_id = f.user_id
		WHERE (u.id = ? OR f.follower_id = ?) 
		AND p.removed <> 1
		ORDER BY 1 DESC`,
		userId, userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var post model.Post
		var user model.User
		if err = rows.Scan(
			&post.UUID,
			&post.Title,
			&post.Body,
			&post.Likes,
			&post.Created,
			&user.ID,
			&user.Name,
			&user.Username,
		); err != nil {
			return nil, err
		}
		post.User = &user
		posts = append(posts, post)
	}

	return posts, nil
}
