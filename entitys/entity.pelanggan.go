package entitys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type EntityPelanggan interface {
	EntityCreate(input *schemas.SchemaPelanggan) (*models.Pelanggan, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaPelanggan) (*models.Pelanggan, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.Pelanggan, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaPelanggan) (*models.Pelanggan, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaPelanggan) (*models.Pelanggan, schemas.SchemaDatabaseError)
}
