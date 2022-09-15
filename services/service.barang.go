package services

import (
	"github.com/jordanlaksono/golang-backend-pos.git/entitys"
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type serviceBarang struct {
	barang entitys.EntityBarang
}

func NewServiceBarang(barang entitys.EntityBarang) *serviceBarang {
	return &serviceBarang{barang: barang}
}

/**
* ==========================================
* Service Create New Outlet Teritory
*===========================================
 */

func (s *serviceBarang) EntityCreate(input *schemas.SchemaBarang) (*models.Barang, schemas.SchemaDatabaseError) {
	var barang schemas.SchemaBarang

	barang.Barang_kode = input.Barang_kode
	barang.Barang_nama = input.Barang_nama
	barang.Barang_kelompok_id = input.Barang_kelompok_id
	barang.Barang_supplier_id = input.Barang_supplier_id
	barang.Barang_keterangan = input.Barang_keterangan
	barang.Barang_image = input.Barang_image

	res, err := s.barang.EntityCreate(&barang)
	return res, err
}

/**
* ==========================================
* Service Result All Outlet Teritory
*===========================================
 */

func (s *serviceBarang) EntityResults() (*[]models.Barang, schemas.SchemaDatabaseError) {
	res, err := s.barang.EntityResults()
	return res, err
}

/**
* ==========================================
* Service Result Outlet By ID Teritory
*===========================================
 */

func (s *serviceBarang) EntityResult(input *schemas.SchemaBarang) (*models.Barang, schemas.SchemaDatabaseError) {
	var barang schemas.SchemaBarang
	barang.Barang_id = input.Barang_id

	res, err := s.barang.EntityResult(&barang)
	return res, err
}

/**
* ==========================================
* Service Delete Outlet By ID Teritory
*===========================================
 */

func (s *serviceBarang) EntityDelete(input *schemas.SchemaBarang) (*models.Barang, schemas.SchemaDatabaseError) {
	var barang schemas.SchemaBarang
	barang.Barang_id = barang.Barang_id

	res, err := s.barang.EntityDelete(&barang)
	return res, err
}

/**
* ==========================================
* Service Update Outlet By ID Teritory
*===========================================
 */

func (s *serviceBarang) EntityUpdate(input *schemas.SchemaBarang) (*models.Barang, schemas.SchemaDatabaseError) {
	var barang schemas.SchemaBarang

	barang.Barang_kode = input.Barang_kode
	barang.Barang_nama = input.Barang_nama
	barang.Barang_kelompok_id = input.Barang_kelompok_id
	barang.Barang_supplier_id = input.Barang_supplier_id
	barang.Barang_keterangan = input.Barang_keterangan
	barang.Barang_image = input.Barang_image

	res, err := s.barang.EntityUpdate(&barang)
	return res, err
}
