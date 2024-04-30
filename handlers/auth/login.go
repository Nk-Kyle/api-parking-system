package auth

import (
	"api-parking-system/payload"
	"api-parking-system/repository"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Login godoc
// @Summary Login to the system
// @Description Login to the system
// @Tags auth
// @Accept  json
// @Produce  json
// @Param body body LoginRequest true "User data"
// @Success 200 {string} string "User logged in successfully"
func Login(c *gin.Context) {
	var body payload.LoginUserRequest
	if c.ShouldBind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	// Check if both email and nik are empty
	if body.Email == "" && body.Nik == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email or NIK must be provided",
		})
		return
	}

	doc, err := repository.GetUserByEmailorNik(body.Email, body.Nik)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(doc.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "api-parking-system",
		Subject:   doc.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error generating token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}
