package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"milkyway/internal/delivery"
	"milkyway/internal/repository"
	"milkyway/internal/services"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepo(db)

	userService := services.NewUserService(userRepo)

	userHandler := delivery.NewUserHandler(userService)

	users := r.Group("api/v1/users")
	{
		users.GET("/", userHandler.GetAllUsers)
		users.POST("/", userHandler.CreateUser)
		users.GET("/:id", userHandler.GetUserById)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE(":id", userHandler.DeleteUser)
	}

	postHandler := delivery.PostHandler{}
	posts := r.Group("api/v1/posts")
	{
		posts.GET("/", postHandler.GetAllPosts)
		posts.POST("/", postHandler.CreatePost)
		posts.GET("/:id", postHandler.GetPostById)
		posts.PUT("/:id", postHandler.UpdatePost)
		posts.DELETE("/:id", postHandler.DeletePost)
	}

	commentHandler := delivery.CommentHandler{}
	comments := r.Group("api/v1/posts")
	{
		comments.POST("/", commentHandler.CreateComment)
		comments.PUT("/:id", commentHandler.UpdateComment)
		comments.DELETE("/:id", commentHandler.DeleteComment)
	}
}
