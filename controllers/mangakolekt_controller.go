package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// mk1.0.0_linux.tar.gz
// mangakolekt1.0.0_linux.tar.gz
// mangakolekt_1.0.0_linux.tar.gz
func GetAllVersions(c *gin.Context) {
	entries, err := os.ReadDir("static/mangakolekt")
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "No manga versions found")
		return
	}
	fmt.Println("Entries")
	fmt.Println(entries)
	for _, e := range entries {
		fmt.Println(e)
	}
	c.String(http.StatusAccepted, "ok")
}

func AddNewVersion(c *gin.Context) {

	if err := c.Request.ParseMultipartForm(32); err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, "Error parsing form data: %v", err)
		return
	}

	name := "mangakolekt_1.0.0_linux.tar.gz"

	fileParts := strings.Split(name, "_")

	version := fileParts[1]
	os := fileParts[2]

}
