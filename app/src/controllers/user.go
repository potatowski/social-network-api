package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"social-api/src/database"
	"social-api/src/model"
	"social-api/src/repository"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user model.User

	if err = json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repository.NewRepositoryUser(db)

	userID, err := userRepository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("user: %d", userID)))
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching all users"))
}

func SearchUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching user by id"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user"))
}
