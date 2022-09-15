package services

import (
	"github.com/jordanlaksono/golang-backend-pos.git/entitys"
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type servicePelanggan struct {
	pelanggan entitys.EntityPelanggan
}

func NewServicePelanggan(pelanggan entitys.EntityPelanggan) *servicePelanggan {
	return &servicePelanggan{pelanggan: pelanggan}
}

func (s servicePelanggan) EntityCreate(input *schemas.SchemaPelanggan) (*models.Pelanggan, schemas.SchemaDatabaseError) {
	//TODO implement me
	var pelanggan schemas.SchemaPelanggan

	pelanggan.Pelanggan_kode = input.Pelanggan_kode
	pelanggan.Pelanggan_nama = input.Pelanggan_nama
	pelanggan.Pelanggan_alamat = input.Pelanggan_alamat
	pelanggan.Pelanggan_no_telp_1 = input.Pelanggan_no_telp_1
	pelanggan.Pelanggan_no_telp_2 = input.Pelanggan_no_telp_2
	pelanggan.Pelanggan_email = input.Pelanggan_email
	pelanggan.Pelanggan_kota = input.Pelanggan_kota
	pelanggan.Pelanggan_kode_pos = input.Pelanggan_kode_pos
	pelanggan.Pelanggan_catatan = input.Pelanggan_catatan
	pelanggan.Pelanggan_top_id = input.Pelanggan_top_id
	pelanggan.Pelanggan_status = input.Pelanggan_status
	pelanggan.Penjualan_saldo = input.Penjualan_saldo

	res, err := s.pelanggan.EntityCreate(&pelanggan)
	return res, err
}

func (s servicePelanggan) EntityResult(input *schemas.SchemaPelanggan) (*models.Pelanggan, schemas.SchemaDatabaseError) {
	//TODO implement me
	var pelanggan schemas.SchemaPelanggan
	pelanggan.Pelanggan_id = input.Pelanggan_id

	res, err := s.pelanggan.EntityResult(&pelanggan)
	return res, err
}

func (s servicePelanggan) EntityResults() (*[]models.Pelanggan, schemas.SchemaDatabaseError) {
	//TODO implement me
	res, err := s.pelanggan.EntityResults()
	return res, err
}

func (s servicePelanggan) EntityDelete(input *schemas.SchemaPelanggan) (*models.Pelanggan, schemas.SchemaDatabaseError) {
	//TODO implement me
	var pelanggan schemas.SchemaPelanggan
	pelanggan.Pelanggan_id = pelanggan.Pelanggan_id

	res, err := s.pelanggan.EntityDelete(&pelanggan)
	return res, err
}

func (s servicePelanggan) EntityUpdate(input *schemas.SchemaPelanggan) (*models.Pelanggan, schemas.SchemaDatabaseError) {
	//TODO implement me
	var pelanggan schemas.SchemaPelanggan

	pelanggan.Pelanggan_kode = input.Pelanggan_kode
	pelanggan.Pelanggan_nama = input.Pelanggan_nama
	pelanggan.Pelanggan_alamat = input.Pelanggan_alamat
	pelanggan.Pelanggan_no_telp_1 = input.Pelanggan_no_telp_1
	pelanggan.Pelanggan_no_telp_2 = input.Pelanggan_no_telp_2
	pelanggan.Pelanggan_email = input.Pelanggan_email
	pelanggan.Pelanggan_kota = input.Pelanggan_kota
	pelanggan.Pelanggan_kode_pos = input.Pelanggan_kode_pos
	pelanggan.Pelanggan_catatan = input.Pelanggan_catatan
	pelanggan.Pelanggan_top_id = input.Pelanggan_top_id
	pelanggan.Pelanggan_status = input.Pelanggan_status
	pelanggan.Penjualan_saldo = input.Penjualan_saldo

	res, err := s.pelanggan.EntityUpdate(&pelanggan)
	return res, err
}
