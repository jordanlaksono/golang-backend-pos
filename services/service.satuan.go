package services

import (
	"github.com/jordanlaksono/golang-backend-pos.git/entitys"
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type serviceSatuan struct {
	satuan entitys.EntitySatuan
}

func NewServiceSatuan(satuan entitys.EntitySatuan) *serviceSatuan {
	return &serviceSatuan{satuan: satuan}
}

/**
* ==========================================
* Service Create New Outlet Teritory
*===========================================
 */

func (s *serviceSatuan) EntityCreate(input *schemas.SchemaSatuan) (*models.Satuan, schemas.SchemaDatabaseError) {
	var satuan schemas.SchemaSatuan
	satuan.Satuan_kode = input.Satuan_kode
	satuan.Satuan_nama = input.Satuan_nama

	res, err := s.satuan.EntityCreate(&satuan)
	return res, err
}

/**
* ==========================================
* Service Result All Outlet Teritory
*===========================================
 */

func (s *serviceSatuan) EntityResults() (*[]models.Satuan, schemas.SchemaDatabaseError) {
	res, err := s.satuan.EntityResults()
	return res, err
}

/**
* ==========================================
* Service Result Outlet By ID Teritory
*===========================================
 */

func (s *serviceSatuan) EntityResult(input *schemas.SchemaSatuan) (*models.Satuan, schemas.SchemaDatabaseError) {
	var satuan schemas.SchemaSatuan
	satuan.Satuan_id = input.Satuan_id

	res, err := s.satuan.EntityResult(&satuan)
	return res, err
}

/**
* ==========================================
* Service Delete Outlet By ID Teritory
*===========================================
 */

func (s *serviceSatuan) EntityDelete(input *schemas.SchemaSatuan) (*models.Satuan, schemas.SchemaDatabaseError) {
	var satuan schemas.SchemaSatuan
	satuan.Satuan_id = satuan.Satuan_id

	res, err := s.satuan.EntityDelete(&satuan)
	return res, err
}

/**
* ==========================================
* Service Update Outlet By ID Teritory
*===========================================
 */

func (s *serviceSatuan) EntityUpdate(input *schemas.SchemaSatuan) (*models.Satuan, schemas.SchemaDatabaseError) {
	var satuan schemas.SchemaSatuan
	satuan.Satuan_kode = input.Satuan_kode
	satuan.Satuan_nama = input.Satuan_nama

	res, err := s.satuan.EntityUpdate(&satuan)
	return res, err
}
