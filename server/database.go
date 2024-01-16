package server

import (
	"fmt"
	"os"

	"github.com/AsentientBanana/admin/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() {
	db, err := gorm.Open(sqlite.Open("site.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Admin{})
	db.AutoMigrate(&models.Project{})
	admin_username := os.Getenv("DEFAULT_ADMIN_USER")
	if admin_username == "" {
		admin_username = "admin"
		// panic("Environment variable for admin not set")
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(admin_username), 10)

	if err != nil {
		fmt.Println("Problem generating user info")
		return
	}

	db.FirstOrCreate(&models.Admin{Username: "admin", Password: string(hashedPass)})
}
