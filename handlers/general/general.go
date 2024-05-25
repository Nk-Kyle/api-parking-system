package general

import (
	"api-parking-system/models"
	"api-parking-system/repository"
	"context"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	var global models.Global
	weekData, err := repository.GetThisWeekGlobal()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get Info",
		})
		return
	}

	dataMap := make(map[string]models.Global)
	for weekData.Next(context.Background()) {
		err := weekData.Decode(&global)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed to get Info",
			})
			return
		}
		date := global.Date.UTC().Format("2006/01/02")
		dataMap[date] = global
	}

	startDate := time.Now().UTC().AddDate(0, 0, -6) // Calculate the start date
	endDate := time.Now().UTC()                     // End date is today
	var dates []string
	var billables []int
	var vehicleCount [4]float64
	var vehicles int
	totalDuration := 0
	minDuration := math.MaxInt64
	maxDuration := 0

	for date := startDate; date.Before(endDate) || date.Equal(endDate); date = date.AddDate(0, 0, 1) {
		dateStr := date.Format("2006/01/02")
		dates = append(dates, dateStr)
		if data, ok := dataMap[dateStr]; ok {
			billables = append(billables, data.Billable)
			vehicleCount[0] += float64(data.Motor)
			vehicleCount[1] += float64(data.Mobil)
			vehicleCount[2] += float64(data.Truk)
			vehicleCount[3] += float64(data.Bus)
			totalDuration += data.TotalDuration
			minDuration = int(math.Min(float64(minDuration), float64(data.MinDuration)))
			maxDuration = int(math.Max(float64(maxDuration), float64(data.MaxDuration)))
		} else {
			// If data for the date is missing, append 0 values
			billables = append(billables, 0)
		}
	}

	for _, count := range vehicleCount {
		vehicles += int(count)
	}

	// Calculate percentages
	var vehiclePercentages []float64
	for _, count := range vehicleCount {
		percentage := count / float64(vehicles) * 100
		vehiclePercentages = append(vehiclePercentages, percentage)
	}

	// If the minDuration is still the initial value, set it to 0
	if minDuration == math.MaxInt64 {
		minDuration = 0
	}

	// If any vehiclePercentages are NaN, set them to 0
	for i, percentage := range vehiclePercentages {
		if math.IsNaN(percentage) {
			vehiclePercentages[i] = 0
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"stats": gin.H{
			"dates":  dates,
			"values": billables,
		},
		"types": gin.H{
			"types":       [4]string{"Motor", "Mobil", "Truk", "Bus"},
			"percentages": vehiclePercentages,
		},
		"numStation":     "1",
		"numVehicle":     strconv.Itoa(vehicles),
		"numTransaction": strconv.Itoa(vehicles),
		"avgDuration":    strconv.FormatFloat(float64(totalDuration)/float64(vehicles), 'f', 2, 64),
		"minDuration":    strconv.Itoa(minDuration),
		"maxDuration":    strconv.Itoa(maxDuration),
	})
}
