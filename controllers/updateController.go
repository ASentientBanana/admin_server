package controllers

import (
	"fmt"
	"net/http"

	"github.com/AsentientBanana/admin/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Update(c *gin.Context) {

	var body struct {
		Username    string
		Password    string
		NewPassword string
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Problem reading request",
		})
		return
	}

	db, err := gorm.Open(sqlite.Open("site.db"), &gorm.Config{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Bad request",
		})
		return
	}

	var admin models.Admin

	db.First(&admin, "Username = ?", body.Username)

	hashErr := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(body.Password))

	if hashErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized, bad password",
		})
		return
	}

	pass_bytes := []byte(body.NewPassword)

	_pass, err := bcrypt.GenerateFromPassword(pass_bytes, 10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Problem hashing new password",
		})
		return
	}

	fmt.Println("Updating")

	admin.Password = string(_pass)
	db.Save(admin)

	c.JSON(http.StatusOK, admin)

}
