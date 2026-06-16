package controllers

import (
	"net/http"

	"github.com/AsentientBanana/admin/dto"
	"github.com/AsentientBanana/admin/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context, db *gorm.DB) {

	authDto := dto.AuthDto{}

	if err := c.BindJSON(&authDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Problem reading request",
		})
		return
	}

	token, err := services.AuthenticateUser(authDto)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
