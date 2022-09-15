package services

import (
	"github.com/jordanlaksono/golang-backend-pos.git/entitys"
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type serviceKelompok struct {
	kelompok entitys.EntityKelompok
}

func NewServiceKelompok(kelompok entitys.EntityKelompok) *serviceKelompok {
	return &serviceKelompok{kelompok: kelompok}
}

/**
* ==========================================
* Service Create New Outlet Teritory
*===========================================
 */

func (s *serviceKelompok) EntityCreate(input *schemas.SchemaKelompok) (*models.Kelompok, schemas.SchemaDatabaseError) {
	var kelompok schemas.SchemaKelompok
	kelompok.Kelompok_kode = input.Kelompok_kode
	kelompok.Kelompok_nama = input.Kelompok_nama

	res, err := s.kelompok.EntityCreate(&kelompok)
	return res, err
}

/**
* ==========================================
* Service Result All Outlet Teritory
*===========================================
 */

func (s *serviceKelompok) EntityResults() (*[]models.Kelompok, schemas.SchemaDatabaseError) {
	res, err := s.kelompok.EntityResults()
	return res, err
}

/**
* ==========================================
* Service Result Outlet By ID Teritory
*===========================================
 */

func (s *serviceKelompok) EntityResult(input *schemas.SchemaKelompok) (*models.Kelompok, schemas.SchemaDatabaseError) {
	var kelompok schemas.SchemaKelompok
	kelompok.Kelompok_id = input.Kelompok_id

	res, err := s.kelompok.EntityResult(&kelompok)
	return res, err
}

/**
* ==========================================
* Service Delete Outlet By ID Teritory
*===========================================
 */

func (s *serviceKelompok) EntityDelete(input *schemas.SchemaKelompok) (*models.Kelompok, schemas.SchemaDatabaseError) {
	var kelompok schemas.SchemaKelompok
	kelompok.Kelompok_id = kelompok.Kelompok_id

	res, err := s.kelompok.EntityDelete(&kelompok)
	return res, err
}

/**
* ==========================================
* Service Update Outlet By ID Teritory
*===========================================
 */

func (s *serviceKelompok) EntityUpdate(input *schemas.SchemaKelompok) (*models.Kelompok, schemas.SchemaDatabaseError) {
	var kelompok schemas.SchemaKelompok
	kelompok.Kelompok_kode = input.Kelompok_kode
	kelompok.Kelompok_nama = input.Kelompok_nama

	res, err := s.kelompok.EntityUpdate(&kelompok)
	return res, err
}
