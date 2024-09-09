package server

import (
	"fmt"

	"github.com/AsentientBanana/admin/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() {
	db, err := gorm.Open(sqlite.Open("site.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Project{})

	if err != nil {
		fmt.Println("Problem generating user info")
		return
	}
}
