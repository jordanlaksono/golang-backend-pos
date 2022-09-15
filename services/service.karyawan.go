package services

import (
	"github.com/jordanlaksono/golang-backend-pos.git/entitys"
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type serviceKaryawan struct {
	karyawan entitys.EntityKaryawan
}

func NewServiceKaryawan(karyawan entitys.EntityKaryawan) *serviceKaryawan {
	return &serviceKaryawan{karyawan: karyawan}
}

func (s *serviceKaryawan) EntityInsert(input *schemas.SchemaKaryawan) (*models.Karyawan, schemas.SchemaDatabaseError) {
	var schema schemas.SchemaKaryawan

	schema.Karyawan_nama = input.Karyawan_nama
	schema.Karyawan_alamat = input.Karyawan_alamat
	schema.Karyawan_kota = input.Karyawan_kota
	schema.Karyawan_negara = input.Karyawan_negara
	schema.Karyawan_kode_pos = input.Karyawan_kode_pos
	schema.Karyawan_no_telp = input.Karyawan_no_telp
	schema.Karyawan_email = input.Karyawan_email

	res, err := s.karyawan.EntityInsert(&schema)
	return res, err
}
