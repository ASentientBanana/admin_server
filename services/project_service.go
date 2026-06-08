package services

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/AsentientBanana/admin/constants"
	"github.com/AsentientBanana/admin/dto"
	"github.com/AsentientBanana/admin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetProjects(db *gorm.DB) ([]models.Project, error) {

	var projects []models.Project

	results := db.Find(&projects)

	if results.Error != nil {
		return nil, results.Error

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
	Error    error
	Status   uint
	Projects []models.Project
}

// note(petar): This is, at least for now going to be unused.
func UpdateProjects(c *gin.Context, db *gorm.DB) UpdateReturn {
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

func ParseFormDataToProject(project *dto.CreateForm, dstProject *models.Project, isUpdate bool) error {

	if project.Image != nil {

		image_file, open_image_err := project.Image.Open()

		if open_image_err == nil {
			defer image_file.Close()
			createdFile, createErr := os.Create(path.Join("static", "projects", project.Image.Filename))
			if createErr != nil {
				if !isUpdate {
					dstProject.Image = constants.DEFAULT_IMAGE
				}
			} else {
				_, copy_error := io.Copy(createdFile, image_file)
				if copy_error == nil {
					dstProject.Image = path.Join("download", "projects", project.Image.Filename)
				} else {
					dstProject.Image = constants.DEFAULT_IMAGE
				}
			}
		}
	} else {
		if !isUpdate {
			dstProject.Image = path.Join("download", "projects", constants.DEFAULT_IMAGE)
		}
	}
	dstProject.Name = project.Name
	dstProject.Live = project.Live
	dstProject.Github = project.Github
	dstProject.Description = project.Description
	dstProject.Stack = project.Stack
	dstProject.Position = project.Position
	return nil
}

func CreateProject(project *dto.CreateForm, db *gorm.DB) ([]models.Project, error) {

	//Create project
	new_project := models.Project{}

	if err := ParseFormDataToProject(project, &new_project, false); err != nil {
		return []models.Project{}, err
	}

	created := db.Create(&new_project)

	if created.Error != nil {
		return nil, created.Error
	}

	projects, get_projects_err := GetProjects(db)

	if get_projects_err != nil {
		return nil, get_projects_err
	}

	return projects, nil
}

func UpdateProject(id string, project *dto.CreateForm, db *gorm.DB) ([]models.Project, error) {

	_id, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	updated_project := models.Project{}

	err = ParseFormDataToProject(project, &updated_project, true)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	tx := db.Model(&models.Project{GormBase: models.GormBase{ID: uint(_id)}}).Updates(&updated_project)

	if tx.Error != nil {
		return nil, tx.Error
	}

	projects, get_projects_err := GetProjects(db)

	if get_projects_err != nil {
		return nil, get_projects_err
	}

	return projects, nil
}

func UpdateProjectPositions(updatePositionDto *dto.UpdateProjectPositionsDto, db *gorm.DB) error {

	// Going with a loop for now since its sqlite so its a local db
	for pos := range updatePositionDto.Positions {
		p := updatePositionDto.Positions[pos]
		tx := db.Model(&models.Project{GormBase: models.GormBase{ID: p.ID}}).Update("position", p.Position)
		if tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}
