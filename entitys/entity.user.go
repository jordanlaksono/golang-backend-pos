package entitys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type EntityUser interface {
	EntityRegister(input *schemas.SchemaUser) (*models.User, schemas.SchemaDatabaseError)
	EntityLogin(input *schemas.SchemaUser) (*models.User, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaUser) (*models.User, schemas.SchemaDatabaseError)
}
