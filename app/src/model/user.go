package model

import (
	"errors"
	"social-api/src/security"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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

	if err := user.format(stage); err != nil {
		return err
	}

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

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("Invalid email")
	}

	if stage == Stage_register && user.Password == "" {
		return errors.New("Password is obrigatory")
	}

	return nil
}

func (user *User) format(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)

	if stage == Stage_register {
		hash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(hash)
	}

	return nil
}
