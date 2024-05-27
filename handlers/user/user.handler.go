package user

import (
	"api-parking-system/repository"
	"context"
	"math"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser(c *gin.Context) {
	email := c.GetString("email")
	res, err := repository.GetThisWeekUser(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get User Data"})
		return
	}

	amountPerDate := make(map[string]int64)
	now := time.Now()

	for i := 0; i < 7; i++ {
		date := now.AddDate(0, 0, -i).Format("2006/01/02")
		amountPerDate[date] = 0
	}

	var images []string
	var route []map[string]string

	time_in := "-"
	time_out := "-"
	minDuration := math.MaxInt64
	maxDuration := 0

	loc, _ := time.LoadLocation("Asia/Jakarta") // load Jakarta timezone

	for res.Next(context.Background()) {
		var result bson.M
		if err := res.Decode(&result); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get User Data"})
			return
		}
		// Check if interface is nil
		if result["invoices"] == nil {
			continue
		}
		invoices := result["invoices"].(primitive.A)
		for _, invoice := range invoices {
			invoiceMap := invoice.(primitive.M)
			invoiceDate := time.Unix(int64(invoiceMap["timestamps"].(primitive.M)["created_at"].(primitive.DateTime)/1000), 0).In(loc).Format("2006/01/02")
			amountPerDate[invoiceDate] += int64(invoiceMap["amount"].(int32))

			duration := int(invoiceMap["duration"].(int32))

			if minDuration > duration {
				minDuration = duration
			}
			if maxDuration < duration {
				maxDuration = duration
			}
		}

		vehicle := result["vehicle"].(primitive.M)
		parking_log := vehicle["parking_log"].(primitive.A)

		if len(parking_log) == 0 {
			continue
		}

		last_ele := parking_log[len(parking_log)-1].(primitive.M)
		placeholder := "https://placehold.jp/6a6d8a/ffffff/800x600.jpg?text=No%20Data"
		image_url := last_ele["image_url"].(string)
		if last_ele["state"].(string) == "in" {
			// Only give last data
			images = append(images, image_url, placeholder)
			time_in = time.Unix(int64(last_ele["timestamps"].(primitive.M)["created_at"].(primitive.DateTime)/1000), 0).In(loc).Format("15:04")
		} else if len(parking_log) > 0 {
			second_to_last_ele := parking_log[len(parking_log)-2].(primitive.M)
			image_url_2 := second_to_last_ele["image_url"].(string)
			images = append(images, image_url, image_url_2)

			time_in = time.Unix(int64(second_to_last_ele["timestamps"].(primitive.M)["created_at"].(primitive.DateTime)/1000), 0).In(loc).Format("15:04")
			time_out = time.Unix(int64(last_ele["timestamps"].(primitive.M)["created_at"].(primitive.DateTime)/1000), 0).In(loc).Format("15:04")
		} else {
			images = append(images, placeholder, placeholder)
		}

		// Slice the last three elements, even if there's only one element
		startIndex := len(parking_log) - 3
		if startIndex < 0 {
			startIndex = 0
		}
		parking_log = parking_log[startIndex:]

		for _, log := range parking_log {
			log := log.(primitive.M)
			ele := map[string]string{
				"title":       "Main Parking Station",
				"description": time.Unix(int64(log["timestamps"].(primitive.M)["created_at"].(primitive.DateTime)/1000), 0).In(loc).Format("2006/01/02"),
				"img":         log["image_url"].(string),
			}
			route = append(route, ele)
		}

		for len(route) < 3 {
			ele := map[string]string{
				"title":       "-",
				"description": "-",
				"img":         "-",
			}
			route = append(route, ele)
		}

	}

	// Check images is empty
	if len(images) == 0 {
		images = append(images, "https://placehold.jp/6a6d8a/ffffff/800x600.jpg?text=No%20Data", "https://placehold.jp/6a6d8a/ffffff/800x600.jpg?text=No%20Data")
	}

	if len(route) == 0 {
		for i := 0; i < 3; i++ {
			ele := map[string]string{
				"title":       "-",
				"description": "-",
				"img":         "-",
			}
			route = append(route, ele)
		}
	}

	// Extract keys (dates) from the map and sort
	var keys []string
	for k := range amountPerDate {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var dates []string
	var amounts []int64
	for _, date := range keys {
		dates = append(dates, date)
		amounts = append(amounts, amountPerDate[date])
	}

	// If the minDuration is still the initial value, set it to 0
	if minDuration == math.MaxInt64 {
		minDuration = 0
	}

	c.JSON(http.StatusOK, gin.H{
		"stats": gin.H{
			"dates":  dates,
			"amount": amounts,
		},
		"image":       images,
		"timeIn":      time_in,
		"timeOut":     time_out,
		"minDuration": strconv.Itoa(minDuration),
		"maxDuration": strconv.Itoa(maxDuration),
		"route":       route,
	})
}

func GetUserImages(c *gin.Context) {

	email := c.GetString("email")
	res, err := repository.GetUserImages(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get User Data"})
		return
	}
	var result bson.M
	var images []string

	for res.Next(context.Background()) {
		if err := res.Decode(&result); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get User Data 2"})
			return
		}
		imagesInt := result["images"].(bson.A)

		for _, image := range imagesInt {
			images = append(images, image.(string))
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"images": images,
	})

}
