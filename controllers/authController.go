package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/AsentientBanana/admin/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {

	var body struct {
		Username string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Problem reading request",
		})
		return
	}

	var admin models.Admin

	db, err := gorm.Open(sqlite.Open("site.db"), &gorm.Config{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Problem connecting to the database",
		})
		return
	}

	//Check user
	if db.First(&admin, "Username = ?", body.Username) == nil {
		fmt.Println("Username: ", body.Username)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Bad request for username",
		})
		return
	}

	//Check password
	if bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(body.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Bad request",
		})
		return
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": admin.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

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
