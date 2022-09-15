package entitys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type EntityTop interface {
	EntityCreate(input *schemas.SchemaTop) (*models.Top, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaTop) (*models.Top, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.Top, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaTop) (*models.Top, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaTop) (*models.Top, schemas.SchemaDatabaseError)
}
