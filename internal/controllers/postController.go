package controllers

import (
	"milkyway/internal/services"
	"net/http"
)

func PostController() {
	http.HandleFunc("/posts", services.GetAllPosts)
	http.HandleFunc("/posts/one", services.GetPostById)
	http.HandleFunc("/posts/add", services.CreatePost)
	http.HandleFunc("/posts/update", services.UpdatePost)
	http.HandleFunc("/posts/delete", services.DeletePost)
}
