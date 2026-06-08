package server

import (
	"github.com/AsentientBanana/admin/controllers"
	"github.com/AsentientBanana/admin/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReqBody struct {
	Username string `json:"username" xml:"username" binding:"required"`
	Password string `json:"password" xml:"password" binding:"required"`
}

func InitRoutes(engine *gin.Engine, db *gorm.DB) {

	//projects
	engine.GET("/projects", NewHandler(db, controllers.GetProjects))
	engine.PUT("/projects/positions", NewHandler(db, controllers.UpdateProjectPositions))
	// engine.PUT("/projects", middleware.Authenticate, NewHandler(db, controllers.UpdateProjects))
	engine.PUT("/projects/:id", NewHandler(db, controllers.UpdateProject))
	engine.POST("/projects", NewHandler(db, controllers.CreateProjects))
	engine.DELETE("/projects/:id", NewHandler(db, controllers.DeleteProjects))

	//resume
	engine.GET("/resume", NewHandler(db, controllers.GetResume))
	engine.POST("/resume", middleware.Authenticate, NewHandler(db, controllers.AddResume))

	// auth
	engine.POST("/login", NewHandler(db, controllers.Login))

	////manga kolekt
	//engine.GET("/mangakolekt/versions", NewHandler(db, controllers.GetAllVersions))
	//
}
