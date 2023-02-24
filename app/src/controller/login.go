package controller

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
)

// Login is responsible for login an user
func Login(w http.ResponseWriter, r *http.Request) {
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
	userSaved, err := userRepository.SearchByEmail(user.Email)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if userSaved.ID == 0 {
		response.Error(w, http.StatusUnauthorized, errors.New("Invalid credentials"))
		return
	}

	if err = security.Check(userSaved.Password, user.Password); err != nil {
		response.Error(w, http.StatusUnauthorized, errors.New("Invalid credentials"))
		return
	}

	token, err := security.CreateToken(userSaved.ID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{"token": token})
}
