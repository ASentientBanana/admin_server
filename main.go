package main

import (
	"fmt"
	"github.com/AsentientBanana/admin/server"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	log.SetOutput(os.Stdout)
	fmt.Println("STARTING")

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

	admin_user := os.Getenv("ADMIN_USER")
	admin_password := os.Getenv("ADMIN_PASSWORD")
	fmt.Println("USER:")
	fmt.Println(admin_user)
	if admin_user == "" || admin_password == "" {
		panic("admin user or password missing in environment.")
	}

	server.InitRoutes(r)

	panic(r.Run(":9898"))
}
