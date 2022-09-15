package entitys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type EntityLokasi interface {
	EntityCreate(input *schemas.SchemaLokasi) (*models.Lokasi, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaLokasi) (*models.Lokasi, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.Lokasi, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaLokasi) (*models.Lokasi, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaLokasi) (*models.Lokasi, schemas.SchemaDatabaseError)
}
