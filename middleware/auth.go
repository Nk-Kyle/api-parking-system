package middleware

import (
	"api-parking-system/utils"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware is a middleware that checks if the request has a valid token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			utils.RespondWithError(c.Writer, http.StatusUnauthorized, "No token provided")
			c.Abort()
			return
		}

		// Get the email from the token
		claims, err := ValidateAndExtractToken(token)
		if err != nil {
			utils.RespondWithError(c.Writer, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		email := fmt.Sprintf("%v", claims.Claims.(jwt.MapClaims)["sub"])
		c.Set("email", email)

		c.Next()
	}
}

func ValidateAndExtractToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
