package vehicle

import (
	"api-parking-system/models"
	"api-parking-system/payload"
	"api-parking-system/repository"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func VehicleAction(c *gin.Context) {
	var body payload.VehicleActionRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	body.PlateNumber = strings.ReplaceAll(body.PlateNumber, " ", "")

	user, err := repository.GetUserByPlateNumber(body.PlateNumber)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	vehicle, err := repository.GetVehicleByPlateNumber(*user, body.PlateNumber)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	state := models.In

	// Get latest ParkingLog
	if len(vehicle.ParkingLog) > 0 {
		parkingLog := vehicle.ParkingLog[len(vehicle.ParkingLog)-1]

		if parkingLog.ImageURL == body.ImageUrl {
			c.JSON(400, gin.H{
				"error": "Image has been processed",
			})
			return
		}

		if parkingLog.State == models.In {
			state = models.Out
		}

	}

	newParkingLog := &models.ParkingLog{
		ImageURL: body.ImageUrl,
		State:    state,
	}
	newParkingLog.CreatedAt = time.Now()
	newParkingLog.UpdatedAt = time.Now()

	// Append new parking log to user vehicle
	_, err = repository.PushParkingCollection(*vehicle, *newParkingLog)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Vehicle action success",
		"state":   state,
	})

}
