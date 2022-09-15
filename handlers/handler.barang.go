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
	"path"
)

type handlerBarang struct {
	barang entitys.EntityBarang
}

func NewHandlerBarang(barang entitys.EntityBarang) *handlerBarang {
	return &handlerBarang{barang: barang}
}

func (h *handlerBarang) HandlerCreate(ctx *gin.Context) {
	var body schemas.SchemaBarang
	file, _ := ctx.FormFile("barang_image")

	body.Barang_image = file.Filename
	body.Barang_kode = ctx.PostForm("barang_kode")
	body.Barang_nama = ctx.PostForm("barang_nama")
	body.Barang_kelompok_id = ctx.PostForm("barang_kelompok_id")
	body.Barang_supplier_id = ctx.PostForm("barang_supplier_id")
	body.Barang_keterangan = ctx.PostForm("barang_keterangan")

	err := ctx.SaveUploadedFile(file, path.Join("images/"+file.Filename))
	fmt.Println(err)

	if err != nil {
		helpers.APIResponse(ctx, "File upload error", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorBarang(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.barang.EntityCreate(&body)

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "Barang name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Add new Barang failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Add new Barang successfully", http.StatusCreated, nil)
}

/**
* ======================================
* Handler Results All Product Teritory
*=======================================
 */

func (h *handlerBarang) HandlerResults(ctx *gin.Context) {
	res, error := h.barang.EntityResults()

	if error.Type == "error_results_01" {
		helpers.APIResponse(ctx, "Barang data not found", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Barang data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Result Product By ID Teritory
*=======================================
 */

func (h *handlerBarang) HandlerResult(ctx *gin.Context) {
	var body schemas.SchemaBarang
	id := ctx.Param("id")
	body.Barang_id = id

	errors, code := ValidatorBarang(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.barang.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Barang data not found for this id %s ", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Barang data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Delete Product By ID Teritory
*=======================================
 */

func (h *handlerBarang) HandlerDelete(ctx *gin.Context) {
	var body schemas.SchemaBarang
	id := ctx.Param("id")
	body.Barang_id = id

	errors, code := ValidatorBarang(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.barang.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Barang data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Barang data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Barang data for this id %s success", id), http.StatusOK, res)
}

/**
* ======================================
* Handler Update Outlet By ID Teritory
*=======================================
 */

func (h *handlerBarang) HandlerUpdate(ctx *gin.Context) {
	var body schemas.SchemaBarang
	file, _ := ctx.FormFile("barang_image")

	body.Barang_image = file.Filename
	body.Barang_kode = ctx.PostForm("barang_kode")
	body.Barang_nama = ctx.PostForm("barang_nama")
	body.Barang_kelompok_id = ctx.PostForm("barang_kelompok_id")
	body.Barang_supplier_id = ctx.PostForm("barang_supplier_id")
	body.Barang_keterangan = ctx.PostForm("barang_keterangan")

	errors, code := ValidatorBarang(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.barang.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Barang data not found for this id "), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Barang data failed for this id %s"), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Barang data success for this id %s"), http.StatusCreated, nil)
}

/**
* ======================================
*  All Validator User Input For Product
*=======================================
 */

func ValidatorBarang(ctx *gin.Context, input schemas.SchemaBarang, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "create" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Barang_kode",
					Message: "Name is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Barang_nama",
					Message: "SKU is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Barang_kelompok_id",
					Message: "Kelompok Id is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "Barang_kelompok_id",
					Message: "Kelompok Id must be uuid",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Barang_supplier_id",
					Message: "Supplier Id is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "Barang_supplier_id",
					Message: "Supplier Id must be uuid",
				},
			},
		}
	}

	if Type == "result" || Type == "delete" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Barang_id",
					Message: "Barang_id is required on param",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "Barang_id",
					Message: "Barang_id must be uuid",
				},
			},
		}
	}

	if Type == "update" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Barang_kode",
					Message: "Name is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Barang_nama",
					Message: "SKU is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Barang_kelompok_id",
					Message: "Kelompok Id is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "Barang_kelompok_id",
					Message: "Kelompok Id must be uuid",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "Barang_supplier_id",
					Message: "Supplier Id is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "uuid",
					Field:   "Barang_supplier_id",
					Message: "Supplier Id must be uuid",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
