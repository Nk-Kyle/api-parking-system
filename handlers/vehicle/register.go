package vehicle

import (
	"api-parking-system/models"
	"api-parking-system/payload"
	"api-parking-system/repository"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RegisterNewVehicle(c *gin.Context) {
	var body payload.RegisterVehicleRequest

	if c.ShouldBind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	body.PlateNumber = strings.ReplaceAll(body.PlateNumber, " ", "")
	// Check if vehicle has been registered
	vehicle, err := repository.GetVehicleByPlateNumber(body.PlateNumber)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Vehicle already registered",
		})
		return
	}

	if vehicle != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Vehicle already registered",
		})
		return
	}

	// Get user by email
	user, err := repository.GetUserByEmail(c.GetString("email"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found",
		})
		return
	}

	// Set vehicle data
	vehicle = &models.Vehicle{
		PlateNumber: body.PlateNumber,
		Type:        body.Type,
		ParkingLog:  []models.ParkingLog{},
	}

	// Add vehicle to user
	user.Vehicles = append(user.Vehicles, *vehicle)

	// Update user data
	_, err = repository.UpdateUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating user data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Vehicle registered successfully",
	})
	return

}
