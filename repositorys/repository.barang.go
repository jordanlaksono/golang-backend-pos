package repositorys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
	"gorm.io/gorm"
	"net/http"
)

type repositoryBarang struct {
	db *gorm.DB
}

func NewRepositoryBarang(db *gorm.DB) *repositoryBarang {
	return &repositoryBarang{db: db}
}

/**
* ==========================================
* Repository Create New Outlet Teritory
*===========================================
 */

func (r *repositoryBarang) EntityCreate(input *schemas.SchemaBarang) (*models.Barang, schemas.SchemaDatabaseError) {
	var barang models.Barang

	barang.Barang_kode = input.Barang_kode
	barang.Barang_nama = input.Barang_nama
	barang.Barang_kelompok_id = input.Barang_kelompok_id
	barang.Barang_supplier_id = input.Barang_supplier_id
	barang.Barang_keterangan = input.Barang_keterangan
	barang.Barang_image = input.Barang_image

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&barang)

	addbarang := db.Debug().Create(&barang).Commit()

	if addbarang.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_03",
		}
		return &barang, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &barang, <-err
}

/**
* ==========================================
* Repository Results All Outlet Teritory
*===========================================
 */

func (r *repositoryBarang) EntityResults() (*[]models.Barang, schemas.SchemaDatabaseError) {
	var barang []models.Barang

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&barang)

	checkBarangName := db.Debug().Preload("Kelompok").Preload("Supplier").Find(&barang)

	if checkBarangName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &barang, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &barang, <-err
}

/**
* ==========================================
* Repository Result Merchant By ID Teritory
*===========================================
 */

func (r *repositoryBarang) EntityResult(input *schemas.SchemaBarang) (*models.Barang, schemas.SchemaDatabaseError) {
	var barang models.Barang
	barang.Barang_id = input.Barang_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&barang)

	checkSupplierName := db.Debug().First(&barang)

	if checkSupplierName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &barang, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &barang, <-err
}

/**
* ==========================================
* Repository Delete Merchant By ID Teritory
*===========================================
 */

func (r *repositoryBarang) EntityDelete(input *schemas.SchemaBarang) (*models.Barang, schemas.SchemaDatabaseError) {
	var barang models.Barang
	barang.Barang_id = input.Barang_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&barang)

	checkSupplierName := db.Debug().First(&barang)

	if checkSupplierName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &barang, <-err
	}

	deletebarang := db.Debug().Delete(&barang)

	if deletebarang.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &barang, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &barang, <-err
}

/**
* ==========================================
* Repository Update Merchant By ID Teritory
*===========================================
 */

func (r *repositoryBarang) EntityUpdate(input *schemas.SchemaBarang) (*models.Barang, schemas.SchemaDatabaseError) {
	var barang models.Barang
	barang.Barang_id = input.Barang_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&barang)

	checkBarangName := db.Debug().First(&barang)

	if checkBarangName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &barang, <-err
	}

	barang.Barang_kode = input.Barang_kode
	barang.Barang_nama = input.Barang_nama
	barang.Barang_kelompok_id = input.Barang_kelompok_id
	barang.Barang_supplier_id = input.Barang_supplier_id
	barang.Barang_keterangan = input.Barang_keterangan
	barang.Barang_image = input.Barang_image

	updatebarang := db.Debug().Updates(&barang)

	if updatebarang.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &barang, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &barang, <-err
}
