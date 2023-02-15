package model

import "time"

// Post represents a publication created by a user
type Post struct {
	UUID    string    `json:"uuid,omitempty"`
	ID      int64     `json:"id,omitempty"`
	Title   string    `json:"title,omitempty"`
	Body    string    `json:"content,omitempty"`
	Likes   int64     `json:"likes"`
	User    *User     `json:"user,omitempty"`
	Created time.Time `json:"created,omitempty"`
}
