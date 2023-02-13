package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"social-api/src/database"
	"social-api/src/model"
	"social-api/src/repository"
	"social-api/src/response"
	"social-api/src/security"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if err = user.Prepare(model.Stage_register); err != nil {
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
	search := strings.ToLower(r.URL.Query().Get("search"))

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repository.NewRepositoryUser(db)

	users, err := userRepository.Search(search)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, users)
}

func SearchUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
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

	user, err := userRepository.SearchById(userId)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if user.ID == 0 {
		response.Error(w, http.StatusNotFound, errors.New("user not found"))
	}

	response.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
	}

	userIDInToken, err := security.ExtractUserIDToken(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDInToken {
		response.Error(w, http.StatusForbidden, errors.New("you can only update your own user"))
		return
	}

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

	if err = user.Prepare(model.Stage_update); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	defer db.Close()

	userRepository := repository.NewRepositoryUser(db)
	if err = userRepository.Update(userID, user); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	userIDInToken, err := security.ExtractUserIDToken(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDInToken {
		response.Error(w, http.StatusForbidden, errors.New("you can only delete your own user"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repository.NewRepositoryUser(db)
	if err = userRepository.Delete(userID); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
