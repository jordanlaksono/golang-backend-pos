package repositorys

import (
	"net/http"

	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
	"gorm.io/gorm"
)

type repositoryRole struct {
	db *gorm.DB
}

func NewRepositoryRole(db *gorm.DB) *repositoryRole {
	return &repositoryRole{db: db}
}

func (r *repositoryRole) EntityInsert(input *schemas.SchemaRole) (*models.Role, schemas.SchemaDatabaseError) {
	var role models.Role

	role.Role_id = input.Role_id
	role.Role_name = input.Role_name

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&role)

	addNewRole := db.Debug().Create(&role).Commit()
	
	if addNewRole.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_register_02",
		}
		return &role, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &role, <-err
}
