package services_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AsentientBanana/admin/controllers"
	"github.com/AsentientBanana/admin/models"
	"github.com/AsentientBanana/admin/server"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Setup testing in mem to avoid writing to disk
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal("failed to open test db:", err)
	}
	db.AutoMigrate(&models.Project{})

	return db
}

func TestProjectFlow(t *testing.T) {
	db := setupTestDB(t)

	projectsToBeCreated := []*models.Project{
		&models.Project{Name: "Test Project", Stack: "ts, js, postgress", Description: "Some test project", Live: "http://google.com", Image: "", Position: 2},
		&models.Project{Name: "Test Project 2", Stack: "ts, js, postgress", Description: "Some test project", Live: "http://google.com", Image: "", Position: 2},
		&models.Project{Name: "Test Project 3", Stack: "ts, js, postgress", Description: "Some test project", Live: "http://google.com", Image: "", Position: 2},
	}
	//Create the projects
	for i, p := range projectsToBeCreated {
		tx := db.Create(p)
		if tx.Error != nil {
			t.Fatalf("Problem creating project %d: %s", i, projectsToBeCreated[i].Name)
			break
		}
		t.Logf("Created project %s \n", p.Name)
	}

	gin.SetMode(gin.TestMode)

	r := gin.New()

	// Making a route in test mode using the actual structure
	r.GET("/projects", server.NewHandler(db, controllers.GetProjects))

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/projects", nil)
	r.ServeHTTP(w, req)

	//Asserts
	if w.Code != http.StatusOK {
		t.Errorf("Expexted %d got %d", http.StatusOK, w.Code)
	}

	var projects map[string][]models.Project
	if err := json.Unmarshal(w.Body.Bytes(), &projects); err != nil {
		t.Fatal("failed to parse response:", err)
	}

	if len(projects["projects"]) != len(projectsToBeCreated) {
		t.Errorf("expected 3 projects got %d", len(projects["projects"]))
	}

	if projects["projects"][0].Name != projectsToBeCreated[0].Name {
		t.Errorf("expected '%s' got '%s'", projectsToBeCreated[0].Name, projects["projects"][0].Name)
	}

}
