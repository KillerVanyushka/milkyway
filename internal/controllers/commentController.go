package controllers

import (
	"milkyway/internal/services"
	"net/http"
)

func CommentController() {
	http.HandleFunc("/comment/add", services.CreateComment)
	http.HandleFunc("/comment/update", services.UpdateComment)
	http.HandleFunc("/comment/delete", services.DeleteComment)
	http.HandleFunc("/comment/pid/1", services.GetCommentsByPostId)
	http.HandleFunc("/comment/uid/2", services.GetCommentsByUserId)
}
