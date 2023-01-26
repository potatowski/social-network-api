package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user"))
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
