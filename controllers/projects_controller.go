package controllers

import (
	"fmt"
	"net/http"

	"github.com/AsentientBanana/admin/dto"
	"github.com/AsentientBanana/admin/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

func GetProjects(c *gin.Context, db *gorm.DB) {
	projects, err := services.GetProjects(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Problem getting projects",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}

func UpdateProjects(c *gin.Context, db *gorm.DB) {

	if err := c.Request.ParseMultipartForm(32); err != nil {
		c.String(http.StatusBadRequest, "Error parsing form data: %v", err)
		return
	}

	updated := services.UpdateProjects(c, db)

	if updated.Error != nil {
		c.String(int(updated.Status), updated.Error.Error())
		return
	}

	projects, get_projects_err := services.GetProjects(db)

	if get_projects_err != nil {

		fmt.Println(get_projects_err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Problem getting projects",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})

}

func UpdateProject(c *gin.Context, db *gorm.DB) {

	id := c.Param("id")
	form := dto.CreateForm{}

	if err := c.ShouldBindWith(&form, binding.FormMultipart); err != nil {
		fmt.Println("BIND ERROR")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Problem reading update request",
		})
		return
	}

	projects, err := services.UpdateProject(id, &form, db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Problem updating project",
		})

	}

	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}

func DeleteProjects(c *gin.Context, db *gorm.DB) {
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

func CreateProjects(c *gin.Context, db *gorm.DB) {
	var form dto.CreateForm

	fmt.Println(c.Request.PostForm)
	c.Request.ParseForm()

	if err := c.ShouldBindWith(&form, binding.FormMultipart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Problem reading request",
		})
		return
	}

	projects, err := services.CreateProject(&form, db)

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

func UpdateProjectPositions(c *gin.Context, db *gorm.DB) {

	updateProjectPositionsDto := dto.UpdateProjectPositionsDto{}

	if err := c.ShouldBindWith(&updateProjectPositionsDto, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Problem reading request",
		})
		return
	}

	error := services.UpdateProjectPositions(&updateProjectPositionsDto, db)

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update the project positions",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Project positions updated"})
}
