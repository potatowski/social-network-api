package model

import "time"

type User struct {
	ID       uint64 `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	Password string
	Created  time.Time `json:"created,omitempty"`
}