package repositorys

import (
	"net/http"

	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/pkg"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
	"gorm.io/gorm"
)

type repositoryUser struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repositoryUser {
	return &repositoryUser{db: db}
}

/**
* ==========================================
* Repository Register Auth Teritory
*===========================================
 */

func (r *repositoryUser) EntityRegister(input *schemas.SchemaUser) (*models.User, schemas.SchemaDatabaseError) {
	var user models.User

	user.User_id = input.User_id
	user.User_username = input.User_username
	user.User_password = input.User_password
	user.UserKaryawanID = input.UserKaryawanID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().First(&user, "user_username = ?", input.User_username)

	if checkEmailExist.RowsAffected > 0 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusConflict,
			Type: "error_register_01",
		}
		return &user, <-err
	}

	addNewUser := db.Debug().Create(&user).Commit()

	if addNewUser.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_register_02",
		}
		return &user, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &user, <-err
}

/**
* ==========================================
* Repository Login Auth Teritory
*===========================================
 */

func (r *repositoryUser) EntityLogin(input *schemas.SchemaUser) (*models.User, schemas.SchemaDatabaseError) {
	var user models.User
	user.User_username = input.User_username
	user.User_password = input.User_password

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().First(&user, "user_username = ?", input.User_username)

	if checkEmailExist.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_login_02",
		}
		return &user, <-err
	}

	checkPasswordMatch := pkg.ComparePassword(user.User_password, input.User_password)

	if checkPasswordMatch != nil {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusBadRequest,
			Type: "error_login_02",
		}
		return &user, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &user, <-err
}

func (r *repositoryUser) EntityResult(input *schemas.SchemaUser) (*models.User, schemas.SchemaDatabaseError) {
	var user models.User
	user.User_id = input.User_id

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&user)

	checkUserId := db.Debug().Preload("Karyawan").First(&user)

	if checkUserId.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &user, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &user, <-err
}

func (r *repositoryUser) EntityResultMenuByUser(input *schemas.SchemaUser) (*[]models.MenuRole, schemas.SchemaDatabaseError) {
	var menurole []models.MenuRole

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&menurole)

	getMenu := db.Raw("SELECT * FROM v_menu_roles where user_id= ?", input.User_id).Scan(&menurole)

	if getMenu.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &menurole, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &menurole, <-err
}

func (r *repositoryUser) EntityResults() (*[]models.User, schemas.SchemaDatabaseError) {
	var user []models.User

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&user)

	checkUser := db.Debug().
		Preload("Karyawan").
		Find(&user)

	if checkUser.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &user, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &user, <-err
}
