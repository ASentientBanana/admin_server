package controllers

import (
	"fmt"
	"net/http"

	"github.com/AsentientBanana/admin/constants"
	"github.com/AsentientBanana/admin/services"
	"github.com/gin-gonic/gin"
)

func GetResume(c *gin.Context) {

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename*="%s"`, "resume.pdf"))
	fmt.Println(constants.DEFAULT_RESUME)
	c.File(constants.DEFAULT_RESUME)

}

func AddResume(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(32); err != nil {
		c.String(http.StatusBadRequest, "Error parsing form data: %v", err)
		return
	}
	file_field := c.Request.MultipartForm.File["file"]

	if len(file_field) == 0 {
		c.String(http.StatusBadRequest, "Error parsing form data")
		return
	}

	err := services.AddResume(file_field[0])

	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing form data: %v", err)
		return
	}
	c.Status(http.StatusOK)
}
