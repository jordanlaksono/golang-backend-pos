package entitys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type EntityMenu interface {
	EntityInsert(input *schemas.SchemaMenu) (*models.Menu, schemas.SchemaDatabaseError)
}
