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

type handlerSupplier struct {
	supplier entitys.EntitySupplier
}

func NewHandlerSupplier(supplier entitys.EntitySupplier) *handlerSupplier {
	return &handlerSupplier{supplier: supplier}
}

/**
* =====================================
* Handler Create New Outlet Teritory
*======================================
 */

func (h *handlerSupplier) HandlerCreate(ctx *gin.Context) {
	var body schemas.SchemaSupplier
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorSupplier(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.supplier.EntityCreate(&body)

	if error.Type == "error_create_03" {
		helpers.APIResponse(ctx, "Create new Top failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Top successfully", http.StatusCreated, nil)
}

/**
* ======================================
* Handler Results All Outlet Teritory
*=======================================
 */

func (h *handlerSupplier) HandlerResults(ctx *gin.Context) {
	res, error := h.supplier.EntityResults()

	if error.Type == "error_results_01" {
		helpers.APIResponse(ctx, "Top data not found", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Top data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Result Outlet By ID Teritory
*=======================================
 */

func (h *handlerSupplier) HandlerResult(ctx *gin.Context) {
	var body schemas.SchemaSupplier
	id := ctx.Param("id")
	body.Supplier_id = id

	errors, code := ValidatorSupplier(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.supplier.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Top data not found for this id %s ", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Top data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Delete Outlet By ID Teritory
*=======================================
 */

func (h *handlerSupplier) HandlerDelete(ctx *gin.Context) {
	var body schemas.SchemaSupplier
	id := ctx.Param("id")
	body.Supplier_id = id

	errors, code := ValidatorSupplier(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.supplier.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Top data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Top Outlet data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Top data for this id %s success", id), http.StatusOK, res)
}

/**
* ======================================
* Handler Update Outlet By ID Teritory
*=======================================
 */

func (h *handlerSupplier) HandlerUpdate(ctx *gin.Context) {
	var body schemas.SchemaSupplier
	//id := ctx.Param("id")
	//body.Satuan_id = input.id

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorSupplier(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.supplier.EntityUpdate(&body)

	//if error.Type == "error_update_01" {
	//	helpers.APIResponse(ctx, fmt.Sprintf("Satuan data not found for this id %s ", id), error.Code, nil)
	//	return
	//}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Top data failed for this id "), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Top data success for this id "), http.StatusCreated, nil)
}

/**
* ======================================
*  All Validator User Input For Outlet
*=======================================
 */

func ValidatorSupplier(ctx *gin.Context, input schemas.SchemaSupplier, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "create" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Supplier_kode",
					Message: "Kode is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Supplier_nama",
					Message: "Name is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Supplier_alamat",
					Message: "Alamat is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Supplier_kota",
					Message: "Kota is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Supplier_kode_pos",
					Message: "Kode Pos is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Supplier_email",
					Message: "Email is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Supplier_no_telp",
					Message: "No Telp is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Supplier_top_id",
					Message: "Top ID is required on body",
				},
			},
		}
	}

	if Type == "result" || Type == "delete" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Supplier_id",
					Message: "ID is required on param",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "Supplier_id",
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
					Field:   "Supplier_kode",
					Message: "Kode is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Supplier_nama",
					Message: "Name is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Supplier_alamat",
					Message: "Alamat is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Supplier_kota",
					Message: "Kota is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Supplier_kode_pos",
					Message: "Kode Pos is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Supplier_email",
					Message: "Email is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Supplier_no_telp",
					Message: "No Telp is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Supplier_top_id",
					Message: "Top ID is required on body",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
