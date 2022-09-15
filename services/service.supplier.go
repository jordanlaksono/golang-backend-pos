package services

import (
	"github.com/jordanlaksono/golang-backend-pos.git/entitys"
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type serviceSupplier struct {
	supplier entitys.EntitySupplier
}

func NewServiceSupplier(supplier entitys.EntitySupplier) *serviceSupplier {
	return &serviceSupplier{supplier: supplier}
}

/**
* ==========================================
* Service Create New Outlet Teritory
*===========================================
 */

func (s *serviceSupplier) EntityCreate(input *schemas.SchemaSupplier) (*models.Supplier, schemas.SchemaDatabaseError) {
	var supplier schemas.SchemaSupplier

	supplier.Supplier_kode = input.Supplier_kode
	supplier.Supplier_nama = input.Supplier_nama
	supplier.Supplier_alamat = input.Supplier_alamat
	supplier.Supplier_kota = input.Supplier_kota
	supplier.Supplier_kode_pos = input.Supplier_kode_pos
	supplier.Supplier_email = input.Supplier_email
	supplier.Supplier_no_telp = input.Supplier_no_telp
	supplier.Supplier_top_id = input.Supplier_top_id

	res, err := s.supplier.EntityCreate(&supplier)
	return res, err
}

/**
* ==========================================
* Service Result All Outlet Teritory
*===========================================
 */

func (s *serviceSupplier) EntityResults() (*[]models.Supplier, schemas.SchemaDatabaseError) {
	res, err := s.supplier.EntityResults()
	return res, err
}

/**
* ==========================================
* Service Result Outlet By ID Teritory
*===========================================
 */

func (s *serviceSupplier) EntityResult(input *schemas.SchemaSupplier) (*models.Supplier, schemas.SchemaDatabaseError) {
	var supplier schemas.SchemaSupplier
	supplier.Supplier_id = input.Supplier_id

	res, err := s.supplier.EntityResult(&supplier)
	return res, err
}

/**
* ==========================================
* Service Delete Outlet By ID Teritory
*===========================================
 */

func (s *serviceSupplier) EntityDelete(input *schemas.SchemaSupplier) (*models.Supplier, schemas.SchemaDatabaseError) {
	var supplier schemas.SchemaSupplier
	supplier.Supplier_id = supplier.Supplier_id

	res, err := s.supplier.EntityDelete(&supplier)
	return res, err
}

/**
* ==========================================
* Service Update Outlet By ID Teritory
*===========================================
 */

func (s *serviceSupplier) EntityUpdate(input *schemas.SchemaSupplier) (*models.Supplier, schemas.SchemaDatabaseError) {
	var supplier schemas.SchemaSupplier

	supplier.Supplier_kode = input.Supplier_kode
	supplier.Supplier_nama = input.Supplier_nama
	supplier.Supplier_alamat = input.Supplier_alamat
	supplier.Supplier_kota = input.Supplier_kota
	supplier.Supplier_kode_pos = input.Supplier_kode_pos
	supplier.Supplier_email = input.Supplier_email
	supplier.Supplier_no_telp = input.Supplier_no_telp
	supplier.Supplier_top_id = input.Supplier_top_id

	res, err := s.supplier.EntityUpdate(&supplier)
	return res, err
}
