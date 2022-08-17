package entitys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type EntityMenuRole interface {
	EntityInsert(input *schemas.SchemaMenuRole) (*models.MenuRole, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.MenuRole, schemas.SchemaDatabaseError)
}
