package handlers

import (
	"fmt"
	"github.com/jordanlaksono/golang-backend-pos.git/entitys"
	"github.com/jordanlaksono/golang-backend-pos.git/helpers"
	"github.com/jordanlaksono/golang-backend-pos.git/pkg"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type handlerSatuan struct {
	satuan entitys.EntitySatuan
}

func NewHandlerSatuan(satuan entitys.EntitySatuan) *handlerSatuan {
	return &handlerSatuan{satuan: satuan}
}

/**
* =====================================
* Handler Create New Outlet Teritory
*======================================
 */

func (h *handlerSatuan) HandlerCreate(ctx *gin.Context) {
	var body schemas.SchemaSatuan
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorSatuan(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.satuan.EntityCreate(&body)

	if error.Type == "error_create_03" {
		helpers.APIResponse(ctx, "Create new Satuan failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Satuan successfully", http.StatusCreated, nil)
}

/**
* ======================================
* Handler Results All Outlet Teritory
*=======================================
 */

func (h *handlerSatuan) HandlerResults(ctx *gin.Context) {
	res, error := h.satuan.EntityResults()

	if error.Type == "error_results_01" {
		helpers.APIResponse(ctx, "Satuan data not found", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Satuan data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Result Outlet By ID Teritory
*=======================================
 */

func (h *handlerSatuan) HandlerResult(ctx *gin.Context) {
	var body schemas.SchemaSatuan
	id := ctx.Param("id")
	body.Satuan_id = id

	errors, code := ValidatorSatuan(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.satuan.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Satuan data not found for this id %s ", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Satuan data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Delete Outlet By ID Teritory
*=======================================
 */

func (h *handlerSatuan) HandlerDelete(ctx *gin.Context) {
	var body schemas.SchemaSatuan
	id := ctx.Param("id")
	body.Satuan_id = id

	errors, code := ValidatorSatuan(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.satuan.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Satuan data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Satuan Outlet data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Satuan data for this id %s success", id), http.StatusOK, res)
}

/**
* ======================================
* Handler Update Outlet By ID Teritory
*=======================================
 */

func (h *handlerSatuan) HandlerUpdate(ctx *gin.Context) {
	var body schemas.SchemaSatuan
	//id := ctx.Param("id")
	//body.Satuan_id = input.id

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorSatuan(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.satuan.EntityUpdate(&body)

	//if error.Type == "error_update_01" {
	//	helpers.APIResponse(ctx, fmt.Sprintf("Satuan data not found for this id %s ", id), error.Code, nil)
	//	return
	//}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Satuan data failed for this id "), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Satuan data success for this id "), http.StatusCreated, nil)
}

/**
* ======================================
*  All Validator User Input For Outlet
*=======================================
 */

func ValidatorSatuan(ctx *gin.Context, input schemas.SchemaSatuan, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "create" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Satuan_kode",
					Message: "Kode is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Satuan_nama",
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
					Field:   "Satuan_id",
					Message: "ID is required on param",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "Satuan_id",
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
					Field:   "Satuan_id",
					Message: "ID is required on param",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "Satuan_id",
					Message: "ID must be uuid",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Satuan_kode",
					Message: "Kode is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Satuan_nama",
					Message: "Satuan is required on body",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
