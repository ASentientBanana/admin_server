package server

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoadUserEnv() (string, string, error) {

	admin_user := os.Getenv("ADMIN_USER")
	admin_password := os.Getenv("ADMIN_PASSWORD")

	if admin_password == "" || admin_user == "" {
		return "", "", errors.New("Problem loading user env")
	}

	return admin_user, admin_password, nil

}

func NewHandler(db *gorm.DB, cb func(*gin.Context, *gorm.DB)) func(*gin.Context) {
	return func(ctx *gin.Context) {
		cb(ctx, db)
	}
}
