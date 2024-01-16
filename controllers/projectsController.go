package controllers

import (
	"net/http"

	"github.com/AsentientBanana/admin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetProjects(c *gin.Context) {
	db, err := gorm.Open(sqlite.Open("site.db"), &gorm.Config{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Bad request",
		})
		return
	}

	var projects []models.Project

	results := db.Find(&projects)

	if results.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No results found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}

func UpdateProjects(c *gin.Context) {

	c.JSON(http.StatusAccepted, gin.H{
		"ok": "ok",
	})
	return

	var body struct {
		project models.Project
	}

	err := c.Bind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": " bad request",
		})
	}

	db, err := gorm.Open(sqlite.Open("site.db"), &gorm.Config{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Bad request",
		})
		return
	}

	var projects []models.Project

	results := db.Find(&projects)

	if results.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No results found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}

func DeleteProjects(c *gin.Context) {

}

func CreateProjects(c *gin.Context) {

}
