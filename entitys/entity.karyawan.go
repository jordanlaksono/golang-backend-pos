package entitys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type EntityKaryawan interface {
	EntityInsert(input *schemas.SchemaKaryawan) (*models.Karyawan, schemas.SchemaDatabaseError)
}
