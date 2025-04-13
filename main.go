package main

import (
	"log"
	"milkyway/internal/controllers"
	"net/http"
)

func main() {
	controllers.UserController()
	controllers.PostController()
	controllers.CommentController()

	server := &http.Server{
		Addr: ":8089",
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
