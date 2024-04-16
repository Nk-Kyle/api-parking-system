package main

import (
	"api-parking-system/gcs"
	"api-parking-system/handlers/images"
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

	// Connect to GCS
	gcs.ConnectStorage()
}

func main() {
	defer mongodb.Client.Disconnect(mongodb.Context)
	defer gcs.StorageClient.Close()

	// Start the server
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API is up and running!",
		})
	})

	router.POST("/upload", images.Upload)

	auth := router.Group("/auth")
	{
		auth.POST("/register", users.Register)
	}

	router.Run(":8080")

}
