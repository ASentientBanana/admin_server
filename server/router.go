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

	//admin
	server.PUT("/projects", middleware.Validate, controllers.UpdateProjects)
	server.POST("/projects/create", middleware.Validate, controllers.CreateProjects)
	server.DELETE("/projects/:id", middleware.Validate, controllers.DeleteProjects)

	// auth
	server.POST("/login", controllers.Login)
}
