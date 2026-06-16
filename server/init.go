package server

import (
	"io"
	"os"

	"github.com/AsentientBanana/admin/middleware"
	"github.com/AsentientBanana/admin/util"
	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine {
	db, err := InitDatabase()

	if err != nil {
		panic("Failed to initialize the database")
	}

	// Just a check if env variables are present
	_, err = util.LoadUserEnv()

	if err != nil {
		panic(err)
	}

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	engine := gin.Default()
	engine.Static("/download", "./static")

	// middleware
	middleware.AddCorsMiddleware(engine)

	//init routes
	InitRoutes(engine, db)

	return engine
}
