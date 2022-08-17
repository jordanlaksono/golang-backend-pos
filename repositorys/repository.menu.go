package repositorys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
	"gorm.io/gorm"
	"net/http"
)

type repositoryMenu struct {
	db *gorm.DB
}

func NewRepositoryMenu(db *gorm.DB) *repositoryMenu {
	return &repositoryMenu{
		db: db,
	}
}

func (r *repositoryMenu) EntityInsert(input *schemas.SchemaMenu) (*models.Menu, schemas.SchemaDatabaseError) {
	var menu models.Menu

	menu.ID = input.ID
	menu.Menu_nama = input.Menu_nama
	menu.Menu_link = input.Menu_link

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&menu)

	addNewMenu := db.Debug().Create(&menu).Commit()

	if addNewMenu.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_register_02",
		}
		return &menu, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &menu, <-err
}

func (r *repositoryMenu) EntityResults() (*[]models.Menu, schemas.SchemaDatabaseError) {
	var menu []models.Menu

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&menu)

	checkMenuRole := db.Debug().Find(&menu)

	if checkMenuRole.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &menu, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &menu, <-err
}
