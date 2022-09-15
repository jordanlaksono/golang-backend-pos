package entitys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type EntityBarang interface {
	EntityCreate(input *schemas.SchemaBarang) (*models.Barang, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaBarang) (*models.Barang, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.Barang, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaBarang) (*models.Barang, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaBarang) (*models.Barang, schemas.SchemaDatabaseError)
}
