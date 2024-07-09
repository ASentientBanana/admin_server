package controllers

import (
	"net/http"

	"github.com/AsentientBanana/admin/models"
	"github.com/AsentientBanana/admin/services"
	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {
	projects, err := services.GetProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Problem getting projects",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}

func UpdateProjects(c *gin.Context) {
	var body struct {
		Projects []models.Project `json:"projects"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

	projects, err := services.UpdateProjects(body.Projects)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}

func DeleteProjects(c *gin.Context) {
	id := c.Param("id")
	projects, err := services.DeleteProject(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Problem deleting project",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}

func CreateProjects(c *gin.Context) {
	var body struct {
		Project models.Project
	}

	if err := c.Bind(&body.Project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Problem reading request",
		})
		return
	}
	projects, err := services.CreateProject(&body.Project)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Problem creating project",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}
