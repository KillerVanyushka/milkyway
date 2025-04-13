package controllers

import (
	"milkyway/internal/services"
	"net/http"
)

func UserController() {
	http.HandleFunc("/users", services.GetAllUsers)
	http.HandleFunc("/users/one", services.GetUserById)
	http.HandleFunc("/users/add", services.CreateUser)
	http.HandleFunc("/users/update", services.UpdateUser)
	http.HandleFunc("/users/delete", services.DeleteUser)
}
