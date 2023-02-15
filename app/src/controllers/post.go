package controllers

import (
	"net/http"
)

// CreatePost creates a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {}

// SearchPosts searches posts by title or content
func SearchPosts(w http.ResponseWriter, r *http.Request) {}

// SearchPostByUuid searches a post by uuid
func SearchPostByUuid(w http.ResponseWriter, r *http.Request) {}

// UpdatePost updates a post
func UpdatePost(w http.ResponseWriter, r *http.Request) {}

// DeletePost deletes a post
func DeletePost(w http.ResponseWriter, r *http.Request) {}
