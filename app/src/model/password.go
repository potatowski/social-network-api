package model

// Password represents request format to change password
type Password struct {
	New string `json:"password"`
	Old string `json:"oldPassword"`
}
