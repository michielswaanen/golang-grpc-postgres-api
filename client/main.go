package main

import (
	"./service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	account := router.Group("/account")
	{
		account.POST("/login", service.Login)
		account.POST("/register", service.Register)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
