package main

import (
	"github.com/gin-gonic/gin"
	"milkyway/internal/routes"
)

func main() {

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
