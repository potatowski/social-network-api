package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"social-api/src/database"
	"social-api/src/model"
	"social-api/src/repository"
	"social-api/src/response"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user model.User
	if err = json.Unmarshal(body, &user); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repository.NewRepositoryUser(db)

	user.ID, err = userRepository.Create(user)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusCreated, user)
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
