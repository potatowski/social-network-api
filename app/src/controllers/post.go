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

	"github.com/gorilla/mux"
)

// CreatePost creates a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := security.ExtractUserIDToken(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post model.Post
	if err = json.Unmarshal(body, &post); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	post.UserID = userID

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	if err = post.Prepare(model.Stage_register); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	postRepository := repository.NewRepositoryPost(db)
	post.ID, err = postRepository.Create(post)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusCreated, post)
}

// SearchPosts searches posts by following users
func SearchPosts(w http.ResponseWriter, r *http.Request) {
	userId, err := security.ExtractUserIDToken(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepository := repository.NewRepositoryPost(db)
	posts, err := postRepository.SearchUserFollowing(userId)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, posts)
}

// SearchPostByUuid searches a post by uuid
func SearchPostByUuid(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postUUID := params["uuid"]

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepository := repository.NewRepositoryPost(db)
	post, err := postRepository.SearchByUuid(postUUID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, post)
}

// UpdatePost updates a post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userId, err := security.ExtractUserIDToken(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	postUUID := params["uuid"]

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepository := repository.NewRepositoryPost(db)
	post, err := postRepository.SearchByUuid(postUUID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if post.UserID != userId {
		response.Error(w, http.StatusForbidden, errors.New("you can't update this post"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var postUpdated model.Post
	if err = json.Unmarshal(body, &postUpdated); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = postUpdated.Prepare(model.Stage_update); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = postRepository.Update(postUUID, postUpdated); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// DeletePost deletes a post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	userId, err := security.ExtractUserIDToken(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	postUUID := params["uuid"]

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepository := repository.NewRepositoryPost(db)
	post, err := postRepository.SearchByUuid(postUUID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if post.UserID != userId {
		response.Error(w, http.StatusForbidden, errors.New("you can't update this post"))
		return
	}

	if err = postRepository.Delete(postUUID); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// SearchPostsByUser searches posts by user
func SearchPostsByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
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

	postRepository := repository.NewRepositoryPost(db)
	posts, err := postRepository.SearchByUser(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, posts)
}
