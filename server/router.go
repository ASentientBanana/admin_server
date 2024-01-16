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
	server.GET("/api/projects", controllers.GetProjects)

	//admin
	server.PUT("/api/projects", middleware.Validate, controllers.UpdateProjects)

	// auth
	server.POST("/admin/api/login", controllers.Login)
	server.POST("/admin/api/register", controllers.Update)

}
