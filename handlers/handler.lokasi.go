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

type handlerLokasi struct {
	lokasi entitys.EntityLokasi
}

func NewHandlerLokasi(lokasi entitys.EntityLokasi) *handlerLokasi {
	return &handlerLokasi{lokasi: lokasi}
}

/**
* =====================================
* Handler Create New Outlet Teritory
*======================================
 */

func (h *handlerLokasi) HandlerCreate(ctx *gin.Context) {
	var body schemas.SchemaLokasi
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorLokasi(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.lokasi.EntityCreate(&body)

	if error.Type == "error_create_03" {
		helpers.APIResponse(ctx, "Create new Lokasi failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Lokasi successfully", http.StatusCreated, nil)
}

/**
* ======================================
* Handler Results All Outlet Teritory
*=======================================
 */

func (h *handlerLokasi) HandlerResults(ctx *gin.Context) {
	res, error := h.lokasi.EntityResults()

	if error.Type == "error_results_01" {
		helpers.APIResponse(ctx, "Lokasi data not found", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Lokasi data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Result Outlet By ID Teritory
*=======================================
 */

func (h *handlerLokasi) HandlerResult(ctx *gin.Context) {
	var body schemas.SchemaLokasi
	id := ctx.Param("id")
	body.Lokasi_id = id

	errors, code := ValidatorLokasi(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.lokasi.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Lokasi data not found for this id %s ", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Lokasi data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Delete Outlet By ID Teritory
*=======================================
 */

func (h *handlerLokasi) HandlerDelete(ctx *gin.Context) {
	var body schemas.SchemaLokasi
	id := ctx.Param("id")
	body.Lokasi_id = id

	errors, code := ValidatorLokasi(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.lokasi.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Lokasi data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Lokasi data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Lokasi data for this id %s success", id), http.StatusOK, res)
}

/**
* ======================================
* Handler Update Outlet By ID Teritory
*=======================================
 */

func (h *handlerLokasi) HandlerUpdate(ctx *gin.Context) {
	var body schemas.SchemaLokasi
	//id := ctx.Param("id")
	//body.Satuan_id = input.id

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorLokasi(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.lokasi.EntityUpdate(&body)

	//if error.Type == "error_update_01" {
	//	helpers.APIResponse(ctx, fmt.Sprintf("Satuan data not found for this id %s ", id), error.Code, nil)
	//	return
	//}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Lokasi data failed for this id "), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Lokasi data success for this id "), http.StatusCreated, nil)
}

/**
* ======================================
*  All Validator User Input For Outlet
*=======================================
 */

func ValidatorLokasi(ctx *gin.Context, input schemas.SchemaLokasi, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "create" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Lokasi_kode",
					Message: "Kode is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Lokasi_nama",
					Message: "Name is required on body",
				},
			},
		}
	}

	if Type == "result" || Type == "delete" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Lokasi_id",
					Message: "ID is required on param",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "Lokasi_id",
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
					Field:   "Lokasi_id",
					Message: "ID is required on param",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "Lokasi_id",
					Message: "ID must be uuid",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Lokasi_kode",
					Message: "Kode is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Lokasi_nama",
					Message: "Satuan is required on body",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
