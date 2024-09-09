package services

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/AsentientBanana/admin/constants"
	"github.com/AsentientBanana/admin/dto"
	"github.com/AsentientBanana/admin/models"
	"github.com/gin-gonic/gin"
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

func extractFormFieldByName(context *gin.Context, name string, index int, project_field *string) {
	key_prefix := fmt.Sprintf("f%d-", index)

	field_slice := context.Request.MultipartForm.Value[key_prefix+name]
	if len(field_slice) == 0 {
		return
	}

	*project_field = field_slice[0]
}

type UpdateReturn struct {
	Error  error
	Status uint
	// Projects []models.Project
}

func UpdateProjects(c *gin.Context) UpdateReturn {
	db, err := gorm.Open(sqlite.Open("site.db"), &gorm.Config{})

	if err != nil {
		return UpdateReturn{Error: err, Status: http.StatusInternalServerError}
	}
	items := []models.Project{}

	// Check items filed
	items_field := c.Request.MultipartForm.Value["items"]
	if len(items_field) == 0 {
		return UpdateReturn{Status: http.StatusBadRequest, Error: errors.New("items field not provided")}
	}

	itemCount, atoi_err := strconv.Atoi(items_field[0])
	if atoi_err != nil {
		return UpdateReturn{Status: http.StatusBadRequest, Error: errors.New("items field invalid value")}
	}
	for i := 0; i < itemCount; i++ {
		_project := models.Project{}
		key_prefix := fmt.Sprintf("f%d-", i)

		//Check position filed
		_position, pos_err := strconv.Atoi(c.Request.MultipartForm.Value[key_prefix+"position"][0])
		if pos_err != nil {
			_project.Position = _position
		}
		// Check id filed

		idValues := c.Request.MultipartForm.Value[key_prefix+"id"]
		if len(idValues) == 0 {
			continue
		}

		_id, id_err := strconv.Atoi(idValues[0])
		if id_err == nil {
			_project.ID = uint(_id)
		}

		extractFormFieldByName(c, "name", i, &_project.Name)
		extractFormFieldByName(c, "github", i, &_project.Github)
		extractFormFieldByName(c, "live", i, &_project.Live)
		extractFormFieldByName(c, "description", i, &_project.Description)
		extractFormFieldByName(c, "stack", i, &_project.Stack)
		extractFormFieldByName(c, "name", i, &_project.Name)
		//image
		image_field := c.Request.MultipartForm.File[key_prefix+"image"]
		//Todo(petar): fix the problem where if no image is supplied no updates ocure

		if image_field == nil || image_field[0] == nil {
			items = append(items, _project)
			continue
		} else {
			image_name := image_field[0].Filename
			path := `static/projects/` + image_name
			f, err := os.Create(path)

			if err != nil {
				items = append(items, _project)
				continue
			}
			content := []byte{}
			uploaded_image, err := image_field[0].Open()
			if err != nil {
				items = append(items, _project)
				os.Remove(path)
				continue
			}
			uploaded_image.Read(content)
			_, write_err := f.Write(content)
			if write_err != nil {
				items = append(items, _project)
				os.Remove(path)
				continue
			}
			_project.Image = path

			items = append(items, _project)
		}

	}

	for _, item := range items {
		db.Table("projects").Where("id = ?", item.ID).Updates(item)
	}

	return UpdateReturn{Error: nil, Status: 0}
}

func CreateProject(project *dto.CreateForm) ([]models.Project, error) {

	//Create project
	new_project := models.Project{}

	image_file, open_image_err := project.Image.Open()
	defer image_file.Close()

	if project.Image != nil || open_image_err != nil {
		createdFile, createErr := os.Create("static/" + project.Image.Filename)
		if createErr != nil {
			new_project.Image = constants.DEFAULT_IMAGE
		} else {

			_, copy_error := io.Copy(createdFile, image_file)
			if copy_error == nil {
				new_project.Image = "static/" + project.Image.Filename
			} else {
				new_project.Image = constants.DEFAULT_IMAGE
				fmt.Println("Error saving image:")
				fmt.Println(copy_error)
			}
		}
	}

	new_project.Name = project.Name
	new_project.Live = project.Live
	new_project.Github = project.Github
	new_project.Description = project.Description
	new_project.Stack = project.Stack
	new_project.Position = project.Position

	db, err := gorm.Open(sqlite.Open("site.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	created := db.Create(&new_project)

	if created.Error != nil {
		return nil, created.Error
	}

	projects, get_projects_err := GetProjects()

	if get_projects_err != nil {
		return nil, get_projects_err
	}

	return projects, nil
}
