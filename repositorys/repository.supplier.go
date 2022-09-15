package repositorys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
	"gorm.io/gorm"
	"net/http"
)

type repositorySupplier struct {
	db *gorm.DB
}

func NewRepositorySupplier(db *gorm.DB) *repositorySupplier {
	return &repositorySupplier{db: db}
}

/**
* ==========================================
* Repository Create New Outlet Teritory
*===========================================
 */

func (r *repositorySupplier) EntityCreate(input *schemas.SchemaSupplier) (*models.Supplier, schemas.SchemaDatabaseError) {
	var supplier models.Supplier

	supplier.Supplier_kode = input.Supplier_kode
	supplier.Supplier_nama = input.Supplier_nama
	supplier.Supplier_alamat = input.Supplier_alamat
	supplier.Supplier_kota = input.Supplier_kota
	supplier.Supplier_kode_pos = input.Supplier_kode_pos
	supplier.Supplier_email = input.Supplier_email
	supplier.Supplier_no_telp = input.Supplier_no_telp
	supplier.Supplier_top_id = input.Supplier_top_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&supplier)

	addsupplier := db.Debug().Create(&supplier).Commit()

	if addsupplier.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_03",
		}
		return &supplier, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &supplier, <-err
}

/**
* ==========================================
* Repository Results All Outlet Teritory
*===========================================
 */

func (r *repositorySupplier) EntityResults() (*[]models.Supplier, schemas.SchemaDatabaseError) {
	var supplier []models.Supplier

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&supplier)

	checkSupplierName := db.Debug().Preload("Top").Find(&supplier)

	if checkSupplierName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &supplier, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &supplier, <-err
}

/**
* ==========================================
* Repository Result Merchant By ID Teritory
*===========================================
 */

func (r *repositorySupplier) EntityResult(input *schemas.SchemaSupplier) (*models.Supplier, schemas.SchemaDatabaseError) {
	var supplier models.Supplier
	supplier.Supplier_id = input.Supplier_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&supplier)

	checkSupplierName := db.Debug().First(&supplier)

	if checkSupplierName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &supplier, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &supplier, <-err
}

/**
* ==========================================
* Repository Delete Merchant By ID Teritory
*===========================================
 */

func (r *repositorySupplier) EntityDelete(input *schemas.SchemaSupplier) (*models.Supplier, schemas.SchemaDatabaseError) {
	var supplier models.Supplier
	supplier.Supplier_id = input.Supplier_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&supplier)

	checkSupplierName := db.Debug().First(&supplier)

	if checkSupplierName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &supplier, <-err
	}

	deletesupplier := db.Debug().Delete(&supplier)

	if deletesupplier.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &supplier, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &supplier, <-err
}

/**
* ==========================================
* Repository Update Merchant By ID Teritory
*===========================================
 */

func (r *repositorySupplier) EntityUpdate(input *schemas.SchemaSupplier) (*models.Supplier, schemas.SchemaDatabaseError) {
	var supplier models.Supplier
	supplier.Supplier_id = input.Supplier_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&supplier)

	checkTopName := db.Debug().First(&supplier)

	if checkTopName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &supplier, <-err
	}

	//top.Top_kode = input.Top_kode
	//top.Top_nama = input.Top_nama
	//top.Top_jumlah = input.Top_jumlah
	supplier.Supplier_kode = input.Supplier_kode
	supplier.Supplier_nama = input.Supplier_nama
	supplier.Supplier_alamat = input.Supplier_alamat
	supplier.Supplier_kota = input.Supplier_kota
	supplier.Supplier_kode_pos = input.Supplier_kode_pos
	supplier.Supplier_email = input.Supplier_email
	supplier.Supplier_no_telp = input.Supplier_no_telp
	supplier.Supplier_top_id = input.Supplier_top_id

	updatesupplier := db.Debug().Updates(&supplier)

	if updatesupplier.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &supplier, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &supplier, <-err
}
