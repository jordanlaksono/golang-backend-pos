package repositorys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
	"gorm.io/gorm"
	"net/http"
)

type repositoryPelanggan struct {
	db *gorm.DB
}

func NewRepositoryPelanggan(db *gorm.DB) *repositoryPelanggan {
	return &repositoryPelanggan{db: db}
}

func (r *repositoryPelanggan) EntityCreate(input *schemas.SchemaPelanggan) (*models.Pelanggan, schemas.SchemaDatabaseError) {
	var pelanggan models.Pelanggan

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

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&pelanggan)

	//checkKode := db.Debug().First("pelanggan_kode = ?", pelanggan.Pelanggan_kode)
	//
	//if checkKode.RowsAffected > 0 {
	//	err <- schemas.SchemaDatabaseError{
	//		Code: http.StatusConflict,
	//		Type: "error_create_01",
	//	}
	//	return &pelanggan, <-err
	//}

	addpelanggan := db.Debug().Create(&pelanggan).Commit()

	if addpelanggan.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_03",
		}
		return &pelanggan, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &pelanggan, <-err
}

func (r *repositoryPelanggan) EntityResult(input *schemas.SchemaPelanggan) (*models.Pelanggan, schemas.SchemaDatabaseError) {
	var pelanggan models.Pelanggan
	pelanggan.Pelanggan_id = input.Pelanggan_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&pelanggan)

	checkPelangganId := db.Debug().First(&pelanggan)

	if checkPelangganId.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &pelanggan, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &pelanggan, <-err
}

func (r *repositoryPelanggan) EntityResults() (*[]models.Pelanggan, schemas.SchemaDatabaseError) {
	var product []models.Pelanggan

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&product)

	checkProduct := db.Debug().Find(&product)

	if checkProduct.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &product, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &product, <-err
}

func (r *repositoryPelanggan) EntityDelete(input *schemas.SchemaPelanggan) (*models.Pelanggan, schemas.SchemaDatabaseError) {
	var pelanggan models.Pelanggan
	pelanggan.Pelanggan_id = input.Pelanggan_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&pelanggan)

	checkProductId := db.Debug().First(&pelanggan)

	if checkProductId.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &pelanggan, <-err
	}

	deletePelanggan := db.Debug().Delete(&pelanggan)

	if deletePelanggan.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &pelanggan, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &pelanggan, <-err
}

func (r *repositoryPelanggan) EntityUpdate(input *schemas.SchemaPelanggan) (*models.Pelanggan, schemas.SchemaDatabaseError) {
	var pelanggan models.Pelanggan
	pelanggan.Pelanggan_id = input.Pelanggan_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&pelanggan)

	checkProductId := db.Debug().First(&pelanggan)

	if checkProductId.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &pelanggan, <-err
	}

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

	updateProduct := db.Debug().Updates(&pelanggan)

	if updateProduct.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &pelanggan, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &pelanggan, <-err
}
