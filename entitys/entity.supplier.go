package entitys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type EntitySupplier interface {
	EntityCreate(input *schemas.SchemaSupplier) (*models.Supplier, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaSupplier) (*models.Supplier, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.Supplier, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaSupplier) (*models.Supplier, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaSupplier) (*models.Supplier, schemas.SchemaDatabaseError)
}
