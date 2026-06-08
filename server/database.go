package server

import (
	"fmt"

	"github.com/AsentientBanana/admin/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("site.db"), &gorm.Config{})

	if err != nil {
		return nil, err

	}

	sqliteDB, err := db.DB()
	if err != nil {
		panic("Problem loading db")

	}

	// sqlite house keeping to avoid locked error
	sqliteDB.SetMaxOpenConns(1)
	sqliteDB.SetMaxIdleConns(1)
	sqliteDB.SetConnMaxLifetime(0)

	db.AutoMigrate(&models.Project{})

	//Creating a project trigger for positions
	db.Exec(`
			CREATE TRIGGER update_position
			AFTER INSERT ON projects
			FOR EACH ROW
			WHEN NEW.position = -1
			BEGIN
				UPDATE projects SET position = (SELECT COUNT(*) FROM projects)+1 WHERE id = NEW.id;
			END;
		`)

	if err != nil {
		fmt.Println("Problem generating user info")
		return nil, err
	}
	return db, nil
}
