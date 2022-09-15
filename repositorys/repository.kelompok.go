package repositorys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
	"gorm.io/gorm"
	"net/http"
)

type repositoryKelompok struct {
	db *gorm.DB
}

func NewRepositoryKelompok(db *gorm.DB) *repositoryKelompok {
	return &repositoryKelompok{db: db}
}

/**
* ==========================================
* Repository Create New Outlet Teritory
*===========================================
 */

func (r *repositoryKelompok) EntityCreate(input *schemas.SchemaKelompok) (*models.Kelompok, schemas.SchemaDatabaseError) {
	var kelompok models.Kelompok

	kelompok.Kelompok_kode = input.Kelompok_kode
	kelompok.Kelompok_nama = input.Kelompok_nama

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&kelompok)

	addkelompok := db.Debug().Create(&kelompok).Commit()

	if addkelompok.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_03",
		}
		return &kelompok, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &kelompok, <-err
}

/**
* ==========================================
* Repository Results All Outlet Teritory
*===========================================
 */

func (r *repositoryKelompok) EntityResults() (*[]models.Kelompok, schemas.SchemaDatabaseError) {
	var kelompok []models.Kelompok

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&kelompok)

	checkKelompokName := db.Debug().Find(&kelompok)

	if checkKelompokName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &kelompok, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &kelompok, <-err
}

/**
* ==========================================
* Repository Result Merchant By ID Teritory
*===========================================
 */

func (r *repositoryKelompok) EntityResult(input *schemas.SchemaKelompok) (*models.Kelompok, schemas.SchemaDatabaseError) {
	var kelompok models.Kelompok
	kelompok.Kelompok_id = input.Kelompok_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&kelompok)

	checkKelompokName := db.Debug().First(&kelompok)

	if checkKelompokName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &kelompok, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &kelompok, <-err
}

/**
* ==========================================
* Repository Delete Merchant By ID Teritory
*===========================================
 */

func (r *repositoryKelompok) EntityDelete(input *schemas.SchemaKelompok) (*models.Kelompok, schemas.SchemaDatabaseError) {
	var kelompok models.Kelompok
	kelompok.Kelompok_id = input.Kelompok_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&kelompok)

	checkKelompokName := db.Debug().First(&kelompok)

	if checkKelompokName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &kelompok, <-err
	}

	deletekelompok := db.Debug().Delete(&kelompok)

	if deletekelompok.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &kelompok, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &kelompok, <-err
}

/**
* ==========================================
* Repository Update Merchant By ID Teritory
*===========================================
 */

func (r *repositoryKelompok) EntityUpdate(input *schemas.SchemaKelompok) (*models.Kelompok, schemas.SchemaDatabaseError) {
	var kelompok models.Kelompok
	kelompok.Kelompok_id = input.Kelompok_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&kelompok)

	checkKelompokName := db.Debug().First(&kelompok)

	if checkKelompokName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &kelompok, <-err
	}

	kelompok.Kelompok_kode = input.Kelompok_kode
	kelompok.Kelompok_nama = input.Kelompok_nama

	updatekelompok := db.Debug().Updates(&kelompok)

	if updatekelompok.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &kelompok, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &kelompok, <-err
}
