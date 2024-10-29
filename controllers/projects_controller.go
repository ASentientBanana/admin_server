package controllers

import (
	"fmt"
	"net/http"

	"github.com/AsentientBanana/admin/dto"
	"github.com/AsentientBanana/admin/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

	if err := c.Request.ParseMultipartForm(32); err != nil {
		c.String(http.StatusBadRequest, "Error parsing form data: %v", err)
		return
	}

	updated := services.UpdateProjects(c)

	if updated.Error != nil {
		c.String(int(updated.Status), updated.Error.Error())
		return
	}

	projects, get_projects_err := services.GetProjects()

	if get_projects_err != nil {

		fmt.Println(get_projects_err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Problem getting projects",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
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
	var form dto.CreateForm

	fmt.Println(c.Request.PostForm)
	c.Request.ParseForm()

	if err := c.ShouldBindWith(&form, binding.FormMultipart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Problem reading request",
		})
		return
	}

	projects, err := services.CreateProject(&form)

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
