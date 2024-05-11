package main

import (
	"api-parking-system/gcs"
	auth_handler "api-parking-system/handlers/auth"
	"api-parking-system/handlers/images"
	user_handler "api-parking-system/handlers/user"
	vehicle_handler "api-parking-system/handlers/vehicle"
	"api-parking-system/middleware"
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
		auth.POST("/register/", auth_handler.Register)
		auth.POST("/login/", auth_handler.Login)
	}

	user := router.Group("/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/profile", user_handler.GetUser)
	}

	vehicle := router.Group("/vehicle")
	vehicle.Use(middleware.AuthMiddleware())
	{
		vehicle.POST("/register/", vehicle_handler.RegisterNewVehicle)
	}

	router.Run(":8080")

}
