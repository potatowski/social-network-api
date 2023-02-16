package model

import (
	"errors"
	"social-api/src/service"
	"strings"
	"time"
)

// Post represents a publication created by a user
type Post struct {
	UUID    string    `json:"uuid,omitempty"`
	ID      uint64    `json:"id,omitempty"`
	Title   string    `json:"title,omitempty"`
	Body    string    `json:"content,omitempty"`
	Likes   uint64    `json:"likes"`
	UserID  uint64    `json:"user_id,omitempty"`
	User    *User     `json:"user,omitempty"`
	Created time.Time `json:"created,omitempty"`
}

// Prepare call functions to check and format post fields
func (post *Post) Prepare(stage string) error {
	if err := post.check(); err != nil {
		return err
	}

	if err := post.format(stage); err != nil {
		return err
	}

	return nil
}

func (post *Post) check() error {
	if post.Title == "" {
		return errors.New("Title is obrigatory")
	}

	if post.Body == "" {
		return errors.New("Body is obrigatory")
	}

	if (post.User == nil) || (post.User.ID == 0) {
		return errors.New("User is obrigatory")
	}

	return nil
}

func (post *Post) format(stage string) error {
	post.Title = strings.TrimSpace(post.Title)
	post.Body = strings.TrimSpace(post.Body)

	if stage == Stage_register {
		post.UUID = service.CreateUUID()
	}
	return nil
}
