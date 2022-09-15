package repositorys

import (
	"fmt"
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
	"gorm.io/gorm"
	"net/http"
)

type repositoryPenjualan struct {
	db *gorm.DB
}

func NewRepositoryPenjualan(db *gorm.DB) *repositoryPenjualan {
	return &repositoryPenjualan{db: db}
}

func (r repositoryPenjualan) EntityPenjualanByKode(input *schemas.SchemaPenjualan) (*models.Penjualan, schemas.SchemaDatabaseError) {
	//TODO implement me
	var penjualan models.Penjualan

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&penjualan)

	checkPenjualan := db.Debug().Find(&penjualan)
	fmt.Println(checkPenjualan)
	if checkPenjualan.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &penjualan, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &penjualan, <-err
}
