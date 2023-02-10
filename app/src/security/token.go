package security

import (
	"social-api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(id uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["user"] = id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString(config.SecretKey)
}
