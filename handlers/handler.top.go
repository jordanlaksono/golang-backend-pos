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

type handlerTop struct {
	top entitys.EntityTop
}

func NewHandlerTop(top entitys.EntityTop) *handlerTop {
	return &handlerTop{top: top}
}

/**
* =====================================
* Handler Create New Outlet Teritory
*======================================
 */

func (h *handlerTop) HandlerCreate(ctx *gin.Context) {
	var body schemas.SchemaTop
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorTop(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.top.EntityCreate(&body)

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

func (h *handlerTop) HandlerResults(ctx *gin.Context) {
	res, error := h.top.EntityResults()

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

func (h *handlerTop) HandlerResult(ctx *gin.Context) {
	var body schemas.SchemaTop
	id := ctx.Param("id")
	body.Top_id = id

	errors, code := ValidatorTop(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.top.EntityResult(&body)

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

func (h *handlerTop) HandlerDelete(ctx *gin.Context) {
	var body schemas.SchemaTop
	id := ctx.Param("id")
	body.Top_id = id

	errors, code := ValidatorTop(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.top.EntityDelete(&body)

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

func (h *handlerTop) HandlerUpdate(ctx *gin.Context) {
	var body schemas.SchemaTop
	//id := ctx.Param("id")
	//body.Satuan_id = input.id

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorTop(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.top.EntityUpdate(&body)

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

func ValidatorTop(ctx *gin.Context, input schemas.SchemaTop, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "create" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Top_kode",
					Message: "Kode is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Top_nama",
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
					Field:   "Top_id",
					Message: "ID is required on param",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "Top_id",
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
					Field:   "Top_id",
					Message: "ID is required on param",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "Top_id",
					Message: "ID must be uuid",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Top_kode",
					Message: "Kode is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Top_nama",
					Message: "Top is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Top_jumlah",
					Message: "Top Jumlah is required on body",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
