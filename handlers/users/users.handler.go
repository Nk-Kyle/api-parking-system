package users

import (
	"api-parking-system/models"
	"api-parking-system/payload"
	"api-parking-system/repository"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param body body RegisterUserRequest true "User data"
// @Success 201 {string} string "User created successfully"
func Register(c *gin.Context) {
	var body payload.RegisterUserRequest
	if c.ShouldBind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	if !isEmailValid(body.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email",
		})
		return
	}

	doc, err := repository.GetUserByEmail(body.Email)
	if doc != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email has been taken!",
		})
		return
	} else if err != nil && err.Error() != "User not found" {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to hash password",
		})
		return
	}

	user := &models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: string(hash),
	}

	_, err = repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}
