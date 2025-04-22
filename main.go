package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"milkyway/internal/models"
	"milkyway/internal/routes"
)

func main() {
	db, err := gorm.Open(postgres.Open("postgresql://postgres-db:mysecretpassword@localhost:5432/postgres?sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connect to the db ", err)
	}

	err1 := db.AutoMigrate(&models.User{})
	if err1 != nil {
		log.Fatal("Error migrate on the db", err)
	}

	err2 := db.AutoMigrate(&models.Post{})
	if err2 != nil {
		log.Fatal("Error migrate on the db", err1)
	}

	err3 := db.AutoMigrate(&models.Comment{})
	if err3 != nil {
		log.Fatal("Error migrate on the db", err2)
	}

	r := gin.Default()

	routes.SetupRoutes(r, db)

	r.Run(":8089")
}
