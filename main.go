package main

import (
	"io"
	"os"
	"time"

	"github.com/AsentientBanana/admin/server"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server.InitDatabase()
	r := gin.Default()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r.Static("/download", "./static")
	//r.StaticFS("/", http.Dir("/static"))

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Change this to your frontend's origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(cors.Default())

	server.InitRoutes(r)

	r.Run(":9898")
}
