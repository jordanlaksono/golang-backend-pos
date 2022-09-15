package entitys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type EntityPenjualan interface {
	EntityPenjualanByKode(input *schemas.SchemaPenjualan) (*models.Penjualan, schemas.SchemaDatabaseError)
}
