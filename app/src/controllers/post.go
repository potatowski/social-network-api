package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"social-api/src/database"
	"social-api/src/model"
	"social-api/src/repository"
	"social-api/src/response"
	"social-api/src/security"
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

// SearchPosts searches posts by title or content
func SearchPosts(w http.ResponseWriter, r *http.Request) {}

// SearchPostByUuid searches a post by uuid
func SearchPostByUuid(w http.ResponseWriter, r *http.Request) {}

// UpdatePost updates a post
func UpdatePost(w http.ResponseWriter, r *http.Request) {}

// DeletePost deletes a post
func DeletePost(w http.ResponseWriter, r *http.Request) {}
