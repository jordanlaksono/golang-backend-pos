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

type handlerMenuRole struct {
	menurole entitys.EntityMenuRole
}

func NewHandlerMenuRole(menurole entitys.EntityMenuRole) *handlerMenuRole {
	return &handlerMenuRole{menurole: menurole}
}

func (h *handlerMenuRole) HandlerInsert(ctx *gin.Context) {
	var body schemas.SchemaMenuRole
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorMenuRole(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.menurole.EntityInsert(&body)

	if error.Type == "error_register_02" {
		helpers.APIResponse(ctx, "Insert new menu role failed", error.Code, nil)

		return
	}

	helpers.APIResponse(ctx, "Insert new menu role success", http.StatusOK, nil)
}

func ValidatorMenuRole(ctx *gin.Context, input schemas.SchemaMenuRole, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	schema = gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "MenuID",
				Message: "Menu id is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Role_id",
				Message: "Role id is required on body",
			},
		},
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}

func (h *handlerMenuRole) HandlerResults(ctx *gin.Context) {
	res, error := h.menurole.EntityResults()

	if error.Type == "error_results_01" {
		helpers.APIResponse(ctx, "Menu Role data not found", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Menu Role data already to use", http.StatusOK, res)
}
