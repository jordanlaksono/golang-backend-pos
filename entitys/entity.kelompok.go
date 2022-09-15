package entitys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type EntityKelompok interface {
	EntityCreate(input *schemas.SchemaKelompok) (*models.Kelompok, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaKelompok) (*models.Kelompok, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.Kelompok, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaKelompok) (*models.Kelompok, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaKelompok) (*models.Kelompok, schemas.SchemaDatabaseError)
}
