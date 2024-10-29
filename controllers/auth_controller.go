package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *gin.Context) {

	var body struct {
		Username string
		Password string
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Problem reading request",
		})
		return
	}

	// Load env vars
	env_secret := os.Getenv("ADMIN_SECRET")
	admin_user := os.Getenv("ADMIN_USER")
	admin_password := os.Getenv("ADMIN_PASSWORD")

	//authorize user
	if admin_user != body.Username || admin_password != body.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Bad request",
		})
		return
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": admin_user,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	var secret string

	if env_secret == "" {
		secret = admin_password + "-" + admin_user
	} else {
		secret = env_secret
	}

	// Sign and get the complete encoded token as a string using the admin credentials
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error signing token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
