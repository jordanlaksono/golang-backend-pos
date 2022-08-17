package repositorys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
	"gorm.io/gorm"
	"net/http"
)

type repositoryMenuRole struct {
	db *gorm.DB
}

func NewRepositoryMenuRole(db *gorm.DB) *repositoryMenuRole {
	return &repositoryMenuRole{
		db: db,
	}
}

func (r *repositoryMenuRole) EntityInsert(input *schemas.SchemaMenuRole) (*models.MenuRole, schemas.SchemaDatabaseError) {
	var menurole models.MenuRole

	menurole.Menu_role_id = input.Menu_role_id
	menurole.MenuID = input.MenuID
	menurole.Role_id = input.Role_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&menurole)

	addNewMenuRole := db.Debug().Create(&menurole).Commit()

	if addNewMenuRole.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_register_02",
		}
		return &menurole, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &menurole, <-err
}

/**
* ==========================================
* Repository Results All Product Teritory
*===========================================
 */

func (r *repositoryMenuRole) EntityResults() (*[]models.MenuRole, schemas.SchemaDatabaseError) {
	var menurole []models.MenuRole

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&menurole)

	checkMenuRole := db.Debug().Find(&menurole)

	if checkMenuRole.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &menurole, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &menurole, <-err
}
