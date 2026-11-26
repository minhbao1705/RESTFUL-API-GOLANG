package main

import (
	"hoc-golang/internal/db"
	"hoc-golang/internal/handlers"
	"hoc-golang/internal/repository"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	if err := db.InitDB(); err != nil {
		log.Fatal("unnable to connect to db")
	}

	log.Println(db.DB)

	r := gin.Default()

	userRepository := repository.NewSQLUserRepository()
	userHander := handlers.NewUserHander(userRepository)

	r.GET("/api/v1/users/:id", userHander.GetUserById)
	r.POST("/api/v1/users", userHander.CreateUser)

	r.Run(":8080")

}
