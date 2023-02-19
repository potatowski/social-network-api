package security

import (
	"errors"
	"net/http"
	"social-api/src/config"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken create a new token
func CreateToken(id uint64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()
	claims["user"] = id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.SecretKey)
}

// ValidateToken validate if the token is valid
func ValidateToken(r *http.Request) error {
	_, err := extractClaims(r)
	if err != nil {
		return err
	}

	return nil
}

// ExtractUserIDToken extract the user id from the token
func ExtractUserIDToken(r *http.Request) (uint64, error) {
	claims, err := extractClaims(r)
	if err != nil {
		return 0, err
	}

	return uint64(claims["user"].(float64)), nil
}

// extractClaims extract the claims from the token
func extractClaims(r *http.Request) (jwt.MapClaims, error) {
	tokenString, err := getToken(r)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, returnKey)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Invalid token")
}

// getToken extract the token from the request
func getToken(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if len(token) == 0 {
		return "", errors.New("Token not found")
	}

	dataToken := strings.Split(token, " ")
	if (len(dataToken) != 2) || (dataToken[0] != "Bearer") {
		return "", errors.New("Invalid type token")
	}

	return dataToken[1], nil
}

func returnKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("Invalid signing method")
	}

	return config.SecretKey, nil
}
