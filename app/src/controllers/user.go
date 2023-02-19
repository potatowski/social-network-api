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

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepare(model.Stage_register); err != nil {
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

// SearchUsers searches users by name or nickname
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

// SearchUserById searches a user by id
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
		return
	}

	response.JSON(w, http.StatusOK, user)
}

// UpdateUser updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
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

// DeleteUser sets a user as deleted
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

// FollowUser follows a user
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followID, err := security.ExtractUserIDToken(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if followID == userID {
		response.Error(w, http.StatusForbidden, errors.New("you can't follow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repository.NewRepositoryUser(db)
	if err = userRepository.Follow(userID, followID); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// UnfollowUser unfollows a user
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followID, err := security.ExtractUserIDToken(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if followID == userID {
		response.Error(w, http.StatusForbidden, errors.New("you can't unfollow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repository.NewRepositoryUser(db)
	if err = userRepository.Unfollow(userID, followID); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// SearchUserFollowers searches a user's followers
func SearchUserFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
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
	followers, err := userRepository.SearchFollowers(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, followers)
}

// SearchUserFollowing searches a user's following
func SearchUserFollowing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
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
	following, err := userRepository.SearchFollowing(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, following)
}

// UpdateUserPassword updates a user's password
func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
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
		response.Error(w, http.StatusForbidden, errors.New("you can only update your own password"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var password model.Password
	if err = json.Unmarshal(body, &password); err != nil {
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
	oldPassword, err := userRepository.GetPasswordByUserId(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.Check(oldPassword, password.Old); err != nil {
		response.Error(w, http.StatusUnauthorized, errors.New("incorrect password"))
		return
	}

	newPass, err := security.Hash(password.New)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = userRepository.UpdatePassword(userID, string(newPass)); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
