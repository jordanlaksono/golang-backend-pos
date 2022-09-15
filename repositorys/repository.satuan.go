package repositorys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
	"gorm.io/gorm"
	"net/http"
)

type repositorySatuan struct {
	db *gorm.DB
}

func NewRepositorySatuan(db *gorm.DB) *repositorySatuan {
	return &repositorySatuan{db: db}
}

/**
* ==========================================
* Repository Create New Outlet Teritory
*===========================================
 */

func (r *repositorySatuan) EntityCreate(input *schemas.SchemaSatuan) (*models.Satuan, schemas.SchemaDatabaseError) {
	var satuan models.Satuan

	satuan.Satuan_kode = input.Satuan_kode
	satuan.Satuan_nama = input.Satuan_nama

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&satuan)

	addsatuan := db.Debug().Create(&satuan).Commit()

	if addsatuan.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_03",
		}
		return &satuan, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &satuan, <-err
}

/**
* ==========================================
* Repository Results All Outlet Teritory
*===========================================
 */

func (r *repositorySatuan) EntityResults() (*[]models.Satuan, schemas.SchemaDatabaseError) {
	var satuan []models.Satuan

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&satuan)

	checkSatuanName := db.Debug().Order("satuan_created_at DESC").Find(&satuan)

	if checkSatuanName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &satuan, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &satuan, <-err
}

/**
* ==========================================
* Repository Result Merchant By ID Teritory
*===========================================
 */

func (r *repositorySatuan) EntityResult(input *schemas.SchemaSatuan) (*models.Satuan, schemas.SchemaDatabaseError) {
	var satuan models.Satuan
	satuan.Satuan_id = input.Satuan_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&satuan)

	checkSatuanName := db.Debug().First(&satuan)

	if checkSatuanName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &satuan, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &satuan, <-err
}

/**
* ==========================================
* Repository Delete Merchant By ID Teritory
*===========================================
 */

func (r *repositorySatuan) EntityDelete(input *schemas.SchemaSatuan) (*models.Satuan, schemas.SchemaDatabaseError) {
	var satuan models.Satuan
	satuan.Satuan_id = input.Satuan_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&satuan)

	checkSatuanName := db.Debug().First(&satuan)

	if checkSatuanName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &satuan, <-err
	}

	deletesatuan := db.Debug().Delete(&satuan)

	if deletesatuan.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &satuan, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &satuan, <-err
}

/**
* ==========================================
* Repository Update Merchant By ID Teritory
*===========================================
 */

func (r *repositorySatuan) EntityUpdate(input *schemas.SchemaSatuan) (*models.Satuan, schemas.SchemaDatabaseError) {
	var satuan models.Satuan
	satuan.Satuan_id = input.Satuan_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&satuan)

	checkSatuanName := db.Debug().First(&satuan)

	if checkSatuanName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &satuan, <-err
	}

	satuan.Satuan_kode = input.Satuan_kode
	satuan.Satuan_nama = input.Satuan_nama

	updatesatuan := db.Debug().Updates(&satuan)

	if updatesatuan.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &satuan, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &satuan, <-err
}
