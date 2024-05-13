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
	var parkingLog models.ParkingLog

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
		parkingLog = vehicle.ParkingLog[len(vehicle.ParkingLog)-1]

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

	if newParkingLog.State == models.Out {
		err = ProcessVehicleOut(*user, parkingLog, *newParkingLog, *vehicle)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"message": "Vehicle action success",
		"state":   state,
	})

}

func ProcessVehicleOut(user models.User, parkingLog models.ParkingLog, newParkingLog models.ParkingLog, vehicle models.Vehicle) error {
	/*
		Precondition: parkingLog.State == models.Out, parkingLog.ImageURL == body.ImageUrl
		Postcondition: Add new invoice to user
	*/

	// Calculate duration
	duration := int(newParkingLog.CreatedAt.Sub(parkingLog.CreatedAt).Minutes())

	// Calculate parking fee based on type
	var cost int
	switch vehicle.Type {
	case "Motor":
		cost = 40
	case "Mobil":
		cost = 80
	case "Truk":
		cost = 150
	case "Bus":
		cost = 200
	}

	// Create a new Invoice
	invoice := &models.Invoice{
		PlateNumber: vehicle.PlateNumber,
		Duration:    int(duration),
		Amount:      int(duration * cost),
		IsPaid:      false,
		PPM:         cost,
		Type:        vehicle.Type,
	}
	invoice.CreatedAt = time.Now()
	invoice.UpdatedAt = time.Now()

	// Append new invoice to user
	_, err := repository.AddInvoice(invoice, user)

	if err != nil {
		return err
	}

	// Get global
	global, err := repository.GetGlobal()
	if err != nil {
		return err
	}
	if global == nil {
		currentDate := time.Now()
		global = &models.Global{
			Date:          time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, time.UTC),
			Billable:      int(duration * cost),
			Transactions:  0,
			TotalDuration: 0,
			MinDuration:   duration,
			MaxDuration:   duration,
		}
	}

	// Update Stats
	switch vehicle.Type {
	case "Motor":
		global.Motor += 1
	case "Mobil":
		global.Mobil += 1
	case "Truk":
		global.Truk += 1
	case "Bus":
		global.Bus += 1
	}

	global.Transactions += 1
	global.TotalDuration += duration

	if global.MinDuration > duration {
		global.MinDuration = duration
	}
	if global.MaxDuration < duration {
		global.MaxDuration = duration
	}

	_, err = repository.UpdateOrCreateGlobal(global)

	if err != nil {
		return err
	}

	return nil
}
