package main

import (
	"api-parking-system/handlers/users"
	"api-parking-system/mongodb"
	"api-parking-system/utils"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	utils.LoadEnv()

	mongodb.ConnectDB()
	mongodb.InitCollections()
}

func main() {
	defer mongodb.Client.Disconnect(mongodb.Context)

	// Start the server
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API is up and running!",
		})
	})

	auth := router.Group("/auth")
	{
		auth.POST("/register", users.Register)
	}

	router.Run(":8080")

}
