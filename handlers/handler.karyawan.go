package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jordanlaksono/golang-backend-pos.git/entitys"
	"github.com/jordanlaksono/golang-backend-pos.git/helpers"
	"github.com/jordanlaksono/golang-backend-pos.git/pkg"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
	gpc "github.com/restuwahyu13/go-playground-converter"
	"net/http"
)

type handlerKaryawan struct {
	karyawan entitys.EntityKaryawan
}

func NewHandlerKaryawan(karyawan entitys.EntityKaryawan) *handlerKaryawan {
	return &handlerKaryawan{karyawan: karyawan}
}

func (h *handlerKaryawan) HandlerInsert(ctx *gin.Context) {
	var body schemas.SchemaKaryawan
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorKaryawan(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.karyawan.EntityInsert(&body)

	if error.Type == "error_register_02" {
		helpers.APIResponse(ctx, "Insert new menu failed", error.Code, nil)

		return
	}

	helpers.APIResponse(ctx, "Insert new menu success", http.StatusOK, nil)
}

func ValidatorKaryawan(ctx *gin.Context, input schemas.SchemaKaryawan, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	schema = gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Karyawan_nama",
				Message: "Nama Karyawan is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Karyawan_alamat",
				Message: "Karyawan alamat is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Karyawan_kota",
				Message: "Karyawan kota is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Karyawan_negara",
				Message: "Karyawan negara is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Karyawan_kode_pos",
				Message: "Karyawan kode pos is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Karyawan_no_telp",
				Message: "Karyawan no telp is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Karyawan_email",
				Message: "Karyawan email is required on body",
			},
		},
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
