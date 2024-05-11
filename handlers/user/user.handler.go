package user

import (
	"api-parking-system/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	// Get user
	c.JSON(http.StatusOK, gin.H{
		"message": utils.GenerateRandomString(10),
	})

}
