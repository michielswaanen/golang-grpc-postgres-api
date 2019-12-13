package main

import (
	"../client/security"
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
	{
		auth := account.Group("/fetch")
		auth.Use(security.AuthenticationRequired("user"))
		{
			auth.GET("/:id", service.Fetch)
		}
	}



	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	} else {
		log.Println("Server is running on port 8080")
	}
}
