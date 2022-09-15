package services

import (
	"github.com/jordanlaksono/golang-backend-pos.git/entitys"
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type serviceTop struct {
	top entitys.EntityTop
}

func NewServiceTop(top entitys.EntityTop) *serviceTop {
	return &serviceTop{top: top}
}

/**
* ==========================================
* Service Create New Outlet Teritory
*===========================================
 */

func (s *serviceTop) EntityCreate(input *schemas.SchemaTop) (*models.Top, schemas.SchemaDatabaseError) {
	var top schemas.SchemaTop
	top.Top_kode = input.Top_kode
	top.Top_nama = input.Top_nama
	top.Top_jumlah = input.Top_jumlah

	res, err := s.top.EntityCreate(&top)
	return res, err
}

/**
* ==========================================
* Service Result All Outlet Teritory
*===========================================
 */

func (s *serviceTop) EntityResults() (*[]models.Top, schemas.SchemaDatabaseError) {
	res, err := s.top.EntityResults()
	return res, err
}

/**
* ==========================================
* Service Result Outlet By ID Teritory
*===========================================
 */

func (s *serviceTop) EntityResult(input *schemas.SchemaTop) (*models.Top, schemas.SchemaDatabaseError) {
	var top schemas.SchemaTop
	top.Top_id = input.Top_id

	res, err := s.top.EntityResult(&top)
	return res, err
}

/**
* ==========================================
* Service Delete Outlet By ID Teritory
*===========================================
 */

func (s *serviceTop) EntityDelete(input *schemas.SchemaTop) (*models.Top, schemas.SchemaDatabaseError) {
	var top schemas.SchemaTop
	top.Top_id = top.Top_id

	res, err := s.top.EntityDelete(&top)
	return res, err
}

/**
* ==========================================
* Service Update Outlet By ID Teritory
*===========================================
 */

func (s *serviceTop) EntityUpdate(input *schemas.SchemaTop) (*models.Top, schemas.SchemaDatabaseError) {
	var top schemas.SchemaTop
	top.Top_kode = input.Top_kode
	top.Top_nama = input.Top_nama
	top.Top_jumlah = input.Top_jumlah

	res, err := s.top.EntityUpdate(&top)
	return res, err
}
