package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewHandler(db *gorm.DB, cb func(*gin.Context, *gorm.DB)) func(*gin.Context) {
	return func(ctx *gin.Context) {
		cb(ctx, db)
	}
}
