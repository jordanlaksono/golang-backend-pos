package entitys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type EntityRole interface {
	EntityInsert(input *schemas.SchemaRole) (*models.Role, schemas.SchemaDatabaseError)
}
