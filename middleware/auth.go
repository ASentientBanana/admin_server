package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authenticate(c *gin.Context) {

	authHeader := c.Request.Header.Get("Authorization")

	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Check token formating to include Bearer
	splitHeader := strings.Split(authHeader, " ")

	if len(splitHeader) < 2 || splitHeader[0] != "Bearer" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	extractedToken := splitHeader[1]

	//parse jwt
	token, err := jwt.Parse(extractedToken, func(token *jwt.Token) (interface {
	}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("ADMIN_SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	//Extract the data
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	//Check if token expired
	if float64(time.Now().Unix()) > claims["exp"].(float64) || !token.Valid {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Next()
}
