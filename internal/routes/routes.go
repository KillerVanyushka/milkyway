package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"milkyway/internal/auth"
	"milkyway/internal/delivery"
	"milkyway/internal/repository"
	"milkyway/internal/services"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {

	users := r.Group("api/v1/auth")
	{
		users.POST("/login", auth.Login)
		users.POST("/register", auth.Register)

	}
	//userRepo := repository.NewUserRepository(db)
	//
	//userService := services.NewUserService(userRepo)
	//
	//userHandler := delivery.NewUserHandler(userService)
	//
	//users := r.Group("api/v1/users")
	//{
	//	users.GET("/", userHandler.GetAllUsers)
	//	users.POST("/", userHandler.CreateUser)
	//	users.GET("/:id", userHandler.GetUserById)
	//	users.PUT("/:id", userHandler.UpdateUser)
	//	users.DELETE(":id", userHandler.DeleteUser)
	//}

	postRepo := repository.NewPostRepository(db)

	postService := services.NewPostService(postRepo)

	postHandler := delivery.NewPostHandler(postService)

	posts := r.Group("api/v1/posts")
	{
		posts.GET("/", postHandler.GetAllPosts)
		posts.POST("/", postHandler.CreatePost)
		posts.GET("/:id", postHandler.GetPostById)
		posts.PUT("/:id", postHandler.UpdatePost)
		posts.DELETE("/:id", postHandler.DeletePost)
	}

	commentRepo := repository.NewCommentRepository(db)

	commentService := services.NewCommentService(commentRepo)

	commentHandler := delivery.NewCommentHandler(commentService)
	comments := r.Group("api/v1/comments")
	{
		comments.POST("/", commentHandler.CreateComment)
		comments.PUT("/:id", commentHandler.UpdateComment)
		comments.DELETE("/:id", commentHandler.DeleteComment)
	}
}
