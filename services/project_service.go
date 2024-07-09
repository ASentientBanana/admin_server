package services

import (
	"fmt"

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

func UpdateProjects(projects []models.Project) ([]models.Project, error) {
	db, err := gorm.Open(sqlite.Open("site.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	update := []models.Project{}
	create := []models.Project{}

	for _, project := range projects {
		if project.ID == 0 {
			create = append(create, project)
			continue
		}
		update = append(update, project)
	}

	for _, v := range create {
		created := db.Create(&v)
		if created.Error != nil {
			continue
		}
	}

	for _, v := range update {
		fmt.Println(v)
		fmt.Println(v.Position)
		db.Model(v).Updates(v)
	}

	_projects, err := GetProjects()

	if err != nil {
		return []models.Project{}, err
	}

	return _projects, nil
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
	created := db.Create(project)
	if created.Error != nil {
		return nil, err
	}

	//get existing projects
	find_result := db.Find(&projects)
	if find_result.Error != nil {
		return nil, err
	}

	return projects, nil
}
