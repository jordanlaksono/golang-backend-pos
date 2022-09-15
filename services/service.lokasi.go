package services

import (
	"github.com/jordanlaksono/golang-backend-pos.git/entitys"
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type serviceLokasi struct {
	lokasi entitys.EntityLokasi
}

func NewServiceLokasi(lokasi entitys.EntityLokasi) *serviceLokasi {
	return &serviceLokasi{lokasi: lokasi}
}

/**
* ==========================================
* Service Create New Outlet Teritory
*===========================================
 */

func (s *serviceLokasi) EntityCreate(input *schemas.SchemaLokasi) (*models.Lokasi, schemas.SchemaDatabaseError) {
	var lokasi schemas.SchemaLokasi
	lokasi.Lokasi_kode = input.Lokasi_kode
	lokasi.Lokasi_nama = input.Lokasi_nama

	res, err := s.lokasi.EntityCreate(&lokasi)
	return res, err
}

/**
* ==========================================
* Service Result All Outlet Teritory
*===========================================
 */

func (s *serviceLokasi) EntityResults() (*[]models.Lokasi, schemas.SchemaDatabaseError) {
	res, err := s.lokasi.EntityResults()
	return res, err
}

/**
* ==========================================
* Service Result Outlet By ID Teritory
*===========================================
 */

func (s *serviceLokasi) EntityResult(input *schemas.SchemaLokasi) (*models.Lokasi, schemas.SchemaDatabaseError) {
	var lokasi schemas.SchemaLokasi
	lokasi.Lokasi_id = input.Lokasi_id

	res, err := s.lokasi.EntityResult(&lokasi)
	return res, err
}

/**
* ==========================================
* Service Delete Outlet By ID Teritory
*===========================================
 */

func (s *serviceLokasi) EntityDelete(input *schemas.SchemaLokasi) (*models.Lokasi, schemas.SchemaDatabaseError) {
	var lokasi schemas.SchemaLokasi
	lokasi.Lokasi_id = lokasi.Lokasi_id

	res, err := s.lokasi.EntityDelete(&lokasi)
	return res, err
}

/**
* ==========================================
* Service Update Outlet By ID Teritory
*===========================================
 */

func (s *serviceLokasi) EntityUpdate(input *schemas.SchemaLokasi) (*models.Lokasi, schemas.SchemaDatabaseError) {
	var lokasi schemas.SchemaLokasi
	lokasi.Lokasi_kode = input.Lokasi_kode
	lokasi.Lokasi_nama = input.Lokasi_nama

	res, err := s.lokasi.EntityUpdate(&lokasi)
	return res, err
}
