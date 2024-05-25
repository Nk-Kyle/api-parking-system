package main

import (
	"api-parking-system/gcs"
	auth_handler "api-parking-system/handlers/auth"
	general_handler "api-parking-system/handlers/general"
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

	// Configure CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization") // Allow Authorization header
	router := gin.Default()
	router.Use(cors.New(config))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API is up and running!",
		})
	})

	router.POST("/upload", images.Upload)
	router.GET("/info", general_handler.Info)

	auth := router.Group("/auth")
	{
		auth.POST("/register/", auth_handler.Register)
		auth.POST("/login/", auth_handler.Login)
	}

	user := router.Group("/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/profile", user_handler.GetUser)
		user.GET("/images", user_handler.GetUserImages)
	}

	vehicle := router.Group("/vehicle")
	{
		vehicle.POST("/register/", middleware.AuthMiddleware(), vehicle_handler.RegisterNewVehicle)
		vehicle.POST("/action/", vehicle_handler.VehicleAction)
	}

	router.Run(":8080")

}
