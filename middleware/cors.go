package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func AddCorsMiddleware(engine *gin.Engine) {
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://petarkocic.net/"}, // Change this to your frontend's origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	engine.Use(cors.Default())
}
