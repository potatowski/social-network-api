package model

import (
	"errors"
	"social-api/src/service"
	"strings"
	"time"
)

// Comment represents a comment in a post created by a user
type Comment struct {
	UUID     string    `json:"uuid,omitempty"`
	ID       uint64    `json:"id,omitempty"`
	Body     string    `json:"body,omitempty"`
	Likes    uint64    `json:"likes"`
	Dislikes uint64    `json:"dislikes"`
	UserID   uint64    `json:"user_id,omitempty"`
	User     *User     `json:"user,omitempty"`
	PostID   uint64    `json:"post_id,omitempty"`
	Post     *Post     `json:"post,omitempty"`
	Created  time.Time `json:"created,omitempty"`
}
