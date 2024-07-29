package server

import (
	"github.com/AsentientBanana/admin/controllers"
	"github.com/AsentientBanana/admin/middleware"
	"github.com/gin-gonic/gin"
)

type ReqBody struct {
	Username string `json:"username" xml:"username" binding:"required"`
	Password string `json:"password" xml:"password" binding:"required"`
}

func InitRoutes(server *gin.Engine) {

	//projects
	server.GET("/projects", controllers.GetProjects)
	server.PUT("/projects", middleware.Authenticate, controllers.UpdateProjects)
	server.POST("/projects", middleware.Authenticate, controllers.CreateProjects)
	server.DELETE("/projects/:id", middleware.Authenticate, controllers.DeleteProjects)

	//resume
	server.GET("/resume", controllers.GetResume)
	server.POST("/resume", controllers.AddResume)

	// auth
	server.POST("/login", controllers.Login)

	//manga kolekt
	server.GET("/mangakolekt/versions", controllers.GetAllVersions)
}
