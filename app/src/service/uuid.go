package service

import "github.com/google/uuid"

func CreateUUID() string {
	return uuid.NewString()
}
