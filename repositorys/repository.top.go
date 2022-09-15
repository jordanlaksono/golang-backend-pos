package repositorys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
	"gorm.io/gorm"
	"net/http"
)

type repositoryTop struct {
	db *gorm.DB
}

func NewRepositoryTop(db *gorm.DB) *repositoryTop {
	return &repositoryTop{db: db}
}

/**
* ==========================================
* Repository Create New Outlet Teritory
*===========================================
 */

func (r *repositoryTop) EntityCreate(input *schemas.SchemaTop) (*models.Top, schemas.SchemaDatabaseError) {
	var top models.Top

	top.Top_kode = input.Top_kode
	top.Top_nama = input.Top_nama
	top.Top_jumlah = input.Top_jumlah

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&top)

	addtop := db.Debug().Create(&top).Commit()

	if addtop.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_03",
		}
		return &top, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &top, <-err
}

/**
* ==========================================
* Repository Results All Outlet Teritory
*===========================================
 */

func (r *repositoryTop) EntityResults() (*[]models.Top, schemas.SchemaDatabaseError) {
	var top []models.Top

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&top)

	checkTopName := db.Debug().Find(&top)

	if checkTopName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &top, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &top, <-err
}

/**
* ==========================================
* Repository Result Merchant By ID Teritory
*===========================================
 */

func (r *repositoryTop) EntityResult(input *schemas.SchemaTop) (*models.Top, schemas.SchemaDatabaseError) {
	var top models.Top
	top.Top_id = input.Top_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&top)

	checkTopName := db.Debug().First(&top)

	if checkTopName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &top, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &top, <-err
}

/**
* ==========================================
* Repository Delete Merchant By ID Teritory
*===========================================
 */

func (r *repositoryTop) EntityDelete(input *schemas.SchemaTop) (*models.Top, schemas.SchemaDatabaseError) {
	var top models.Top
	top.Top_id = input.Top_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&top)

	checkTopName := db.Debug().First(&top)

	if checkTopName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &top, <-err
	}

	deletesatuan := db.Debug().Delete(&top)

	if deletesatuan.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &top, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &top, <-err
}

/**
* ==========================================
* Repository Update Merchant By ID Teritory
*===========================================
 */

func (r *repositoryTop) EntityUpdate(input *schemas.SchemaTop) (*models.Top, schemas.SchemaDatabaseError) {
	var top models.Top
	top.Top_id = input.Top_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&top)

	checkTopName := db.Debug().First(&top)

	if checkTopName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &top, <-err
	}

	top.Top_kode = input.Top_kode
	top.Top_nama = input.Top_nama
	top.Top_jumlah = input.Top_jumlah

	updatesatuan := db.Debug().Updates(&top)

	if updatesatuan.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &top, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &top, <-err
}
