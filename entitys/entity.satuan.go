package entitys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type EntitySatuan interface {
	EntityCreate(input *schemas.SchemaSatuan) (*models.Satuan, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaSatuan) (*models.Satuan, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.Satuan, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaSatuan) (*models.Satuan, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaSatuan) (*models.Satuan, schemas.SchemaDatabaseError)
}
