package repositorys

import (
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
	"gorm.io/gorm"
	"net/http"
)

type repositoryKaryawan struct {
	db *gorm.DB
}

func NewRepositoryKaryawan(db *gorm.DB) *repositoryKaryawan {
	return &repositoryKaryawan{
		db: db,
	}
}

func (r *repositoryKaryawan) EntityInsert(input *schemas.SchemaKaryawan) (*models.Karyawan, schemas.SchemaDatabaseError) {
	var karyawan models.Karyawan

	karyawan.ID = input.ID
	karyawan.Karyawan_nama = input.Karyawan_nama
	karyawan.Karyawan_alamat = input.Karyawan_alamat
	karyawan.Karyawan_kota = input.Karyawan_kota
	karyawan.Karyawan_negara = input.Karyawan_negara
	karyawan.Karyawan_kode_pos = input.Karyawan_kode_pos
	karyawan.Karyawan_no_telp = input.Karyawan_no_telp
	karyawan.Karyawan_email = input.Karyawan_email

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&karyawan)

	addNewMenu := db.Debug().Create(&karyawan).Commit()

	if addNewMenu.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_register_02",
		}
		return &karyawan, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &karyawan, <-err
}
