package repositorys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
	"gorm.io/gorm"
	"net/http"
)

type repositoryLokasi struct {
	db *gorm.DB
}

func NewRepositoryLokasi(db *gorm.DB) *repositoryLokasi {
	return &repositoryLokasi{db: db}
}

/**
* ==========================================
* Repository Register Auth Teritory
*===========================================
 */

func (r *repositoryLokasi) EntityCreate(input *schemas.SchemaLokasi) (*models.Lokasi, schemas.SchemaDatabaseError) {
	var lokasi models.Lokasi

	lokasi.Lokasi_id = input.Lokasi_id
	lokasi.Lokasi_kode = input.Lokasi_kode
	lokasi.Lokasi_nama = input.Lokasi_nama

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&lokasi)

	checkKodeExist := db.Debug().First(&lokasi, "lokasi_kode = ?", input.Lokasi_nama)

	if checkKodeExist.RowsAffected > 0 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusConflict,
			Type: "error_register_01",
		}
		return &lokasi, <-err
	}

	addNewLokasi := db.Debug().Create(&lokasi).Commit()

	if addNewLokasi.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_register_02",
		}
		return &lokasi, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &lokasi, <-err
}

/**
* ==========================================
* Repository Results All Outlet Teritory
*===========================================
 */

func (r *repositoryLokasi) EntityResults() (*[]models.Lokasi, schemas.SchemaDatabaseError) {
	var lokasi []models.Lokasi

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&lokasi)

	checkLokasiName := db.Debug().Find(&lokasi)

	if checkLokasiName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &lokasi, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &lokasi, <-err
}

/**
* ==========================================
* Repository Result Merchant By ID Teritory
*===========================================
 */

func (r *repositoryLokasi) EntityResult(input *schemas.SchemaLokasi) (*models.Lokasi, schemas.SchemaDatabaseError) {
	var lokasi models.Lokasi
	lokasi.Lokasi_id = input.Lokasi_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&lokasi)

	checkLokasiName := db.Debug().First(&lokasi)

	if checkLokasiName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &lokasi, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &lokasi, <-err
}

/**
* ==========================================
* Repository Delete Merchant By ID Teritory
*===========================================
 */

func (r *repositoryLokasi) EntityDelete(input *schemas.SchemaLokasi) (*models.Lokasi, schemas.SchemaDatabaseError) {
	var lokasi models.Lokasi
	lokasi.Lokasi_id = input.Lokasi_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&lokasi)

	checkLokasiName := db.Debug().First(&lokasi)

	if checkLokasiName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &lokasi, <-err
	}

	deletelokasi := db.Debug().Delete(&lokasi)

	if deletelokasi.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &lokasi, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &lokasi, <-err
}

/**
* ==========================================
* Repository Update Merchant By ID Teritory
*===========================================
 */

func (r *repositoryLokasi) EntityUpdate(input *schemas.SchemaLokasi) (*models.Lokasi, schemas.SchemaDatabaseError) {
	var lokasi models.Lokasi
	lokasi.Lokasi_id = input.Lokasi_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&lokasi)

	checkLokasiName := db.Debug().First(&lokasi)

	if checkLokasiName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &lokasi, <-err
	}

	lokasi.Lokasi_kode = input.Lokasi_kode
	lokasi.Lokasi_nama = input.Lokasi_nama

	updatelokasi := db.Debug().Updates(&lokasi)

	if updatelokasi.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &lokasi, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &lokasi, <-err
}
