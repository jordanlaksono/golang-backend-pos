package services

import (
	"github.com/jordanlaksono/golang-backend-pos.git/entitys"
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type servicePenjualan struct {
	penjualan entitys.EntityPenjualan
}

func NewServicePenjualan(penjualan entitys.EntityPenjualan) *servicePenjualan {
	return &servicePenjualan{penjualan: penjualan}
}

func (s servicePenjualan) EntityPenjualanByKode(input *schemas.SchemaPenjualan) (*models.Penjualan, schemas.SchemaDatabaseError) {
	//TODO implement me
	var penjualan schemas.SchemaPenjualan
	res, err := s.penjualan.EntityPenjualanByKode(&penjualan)
	return res, err
}
