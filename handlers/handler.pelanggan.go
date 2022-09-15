package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jordanlaksono/golang-backend-pos.git/entitys"
	"github.com/jordanlaksono/golang-backend-pos.git/helpers"
	"github.com/jordanlaksono/golang-backend-pos.git/pkg"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
	gpc "github.com/restuwahyu13/go-playground-converter"
	"net/http"
)

type handlerPelanggan struct {
	pelanggan entitys.EntityPelanggan
}

func NewHandlerPelanggan(pelanggan entitys.EntityPelanggan) *handlerPelanggan {
	return &handlerPelanggan{pelanggan: pelanggan}
}

func (h *handlerPelanggan) HandlerCreate(ctx *gin.Context) {
	var body schemas.SchemaPelanggan
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorPelanggan(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.pelanggan.EntityCreate(&body)

	if error.Type == "error_create_03" {
		helpers.APIResponse(ctx, "Create new Pelanggan failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Pelanggan successfully", http.StatusCreated, nil)
}

/**
* ======================================
* Handler Results All Outlet Teritory
*=======================================
 */

func (h *handlerPelanggan) HandlerResults(ctx *gin.Context) {
	res, error := h.pelanggan.EntityResults()

	if error.Type == "error_results_01" {
		helpers.APIResponse(ctx, "Pelanggan data not found", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Pelanggan data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Result Outlet By ID Teritory
*=======================================
 */

func (h *handlerPelanggan) HandlerResult(ctx *gin.Context) {
	var body schemas.SchemaPelanggan
	id := ctx.Param("id")
	body.Pelanggan_id = id

	errors, code := ValidatorPelanggan(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.pelanggan.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Pelanggan data not found for this id %s ", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Pelanggan data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Delete Outlet By ID Teritory
*=======================================
 */

func (h *handlerPelanggan) HandlerDelete(ctx *gin.Context) {
	var body schemas.SchemaPelanggan
	id := ctx.Param("id")
	body.Pelanggan_id = id

	errors, code := ValidatorPelanggan(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.pelanggan.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Pelanggan data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Pelanggan Outlet data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Pelanggan data for this id %s success", id), http.StatusOK, res)
}

/**
* ======================================
* Handler Update Outlet By ID Teritory
*=======================================
 */

func (h *handlerPelanggan) HandlerUpdate(ctx *gin.Context) {
	var body schemas.SchemaPelanggan
	//id := ctx.Param("id")
	//body.Satuan_id = input.id

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorPelanggan(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.pelanggan.EntityUpdate(&body)

	//if error.Type == "error_update_01" {
	//	helpers.APIResponse(ctx, fmt.Sprintf("Satuan data not found for this id %s ", id), error.Code, nil)
	//	return
	//}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Pelanggan data failed for this id "), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Pelanggan data success for this id "), http.StatusCreated, nil)
}

/**
* ======================================
*  All Validator User Input For Outlet
*=======================================
 */

func ValidatorPelanggan(ctx *gin.Context, input schemas.SchemaPelanggan, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "create" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_kode",
					Message: "Kode is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_nama",
					Message: "Name is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_alamat",
					Message: "Alamat is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_no_telp_1",
					Message: "No Telp 1 is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_no_telp_2",
					Message: "No Telp 2 is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_email",
					Message: "Email is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_kota",
					Message: "Kota is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_kode_pos",
					Message: "Kode Pos is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_catatan",
					Message: "Catatn is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_top_id",
					Message: "Top is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_status",
					Message: "Status is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_saldo",
					Message: "Saldo is required on body",
				},
			},
		}
	}

	if Type == "result" || Type == "delete" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_id",
					Message: "ID is required on param",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "Pelanggan_id",
					Message: "ID must be uuid",
				},
			},
		}
	}

	if Type == "update" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_kode",
					Message: "Kode is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_nama",
					Message: "Name is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_alamat",
					Message: "Alamat is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_no_telp_1",
					Message: "No Telp 1 is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_no_telp_2",
					Message: "No Telp 2 is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_email",
					Message: "Email is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_kota",
					Message: "Kota is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_kode_pos",
					Message: "Kode Pos is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_catatan",
					Message: "Catatn is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_top_id",
					Message: "Top is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_status",
					Message: "Status is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Pelanggan_saldo",
					Message: "Saldo is required on body",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
