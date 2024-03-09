package services

import (
	"github.com/AsentientBanana/admin/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetProjects() ([]models.Project, error) {
	db, err := gorm.Open(sqlite.Open("site.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	var projects []models.Project

	results := db.Find(&projects)

	if results.Error != nil {
		return nil, err

	}
	return projects, nil
}

func DeleteProject(id string) ([]models.Project, error) {

	db, err := gorm.Open(sqlite.Open("site.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	result := db.Delete(&models.Project{}, id)
	if result.Error != nil {
		return nil, result.Error
	}

	var projects []models.Project

	db.Find(&projects)

	return projects, nil
}

func CreateProject(project *models.Project) ([]models.Project, error) {

	var projects []models.Project
	db, err := gorm.Open(sqlite.Open("site.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	//TODO: Add validation?
	//Check desired project position
	result := db.Find(&projects)
	if result.Error != nil {
		return nil, err

	}
	//Create project
	create_result := db.Create(project)
	if create_result.Error != nil {
		return nil, err
	}

	//get existing projects
	find_result := db.Find(&projects)
	if find_result.Error != nil {
		return nil, err
	}

	return projects, nil
}
