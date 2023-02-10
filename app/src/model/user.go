package model

import (
	"errors"
	"strings"
	"time"
)

const (
	Stage_register = "REGISTER"
	Stage_update   = "UPDATE"
)

type User struct {
	ID       uint64 `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string
	Removed  bool      `json:"removed,omitempty"`
	Created  time.Time `json:"created,omitempty"`
}

// Prepare call functions to check and format user fields
func (user *User) Prepare(stage string) error {
	if err := user.check(stage); err != nil {
		return err
	}

	user.format()
	return nil
}

func (user *User) check(stage string) error {
	if user.Name == "" {
		return errors.New("Name is obrigatory")
	}

	if user.Username == "" {
		return errors.New("Username is obrigatory")
	}

	if user.Email == "" {
		return errors.New("Email is obrigatory")
	}

	if stage == Stage_register && user.Password == "" {
		return errors.New("Password is obrigatory")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)
}
