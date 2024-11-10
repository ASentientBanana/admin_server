package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/AsentientBanana/admin/services"
	"github.com/gin-gonic/gin"
)

// name format
// mangakolekt1.0.0_linux.tar.gz
func GetAllVersions(c *gin.Context) {
	contents, err := os.ReadDir("static/mangakolekt")
	if err != nil {
		c.String(http.StatusInternalServerError, "No manga versions found")
		return
	}

	versions := make(map[string][]services.VersionEntry)

	for _, content := range contents {

		if !content.IsDir() {
			continue
		}
		osName := content.Name()
		_, ok := versions[osName]

		if !ok {
			versions[osName] = []services.VersionEntry{}
		}
		entryPath := path.Join("static/mangakolekt", osName)
		entryPathAlias := path.Join("download/mangakolekt", osName)
		entries, err := services.GetDirEntries(entryPath, entryPathAlias)
		if err != nil {
			continue
		}

		versions[osName] = append(versions[osName], entries...)
		// for i := 0; i < len(entries); i++ {
		// 	fmt.Println(entries[i])
		// 	versions[osName] = append(versions[osName], entries[i])

		// }
	}

	c.JSON(http.StatusAccepted, gin.H{
		"versions": versions,
	})
}

func AddNewVersion(c *gin.Context) {

	if err := c.Request.ParseMultipartForm(32); err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, "Error parsing form data: %v", err)
		return
	}

	// name := "mangakolekt_1.0.0_linux.tar.gz"

	// fileParts := strings.Split(name, "_")

	// version := fileParts[1]
	// fileParts[2]

}
