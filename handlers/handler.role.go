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

type handlerRole struct {
	role entitys.EntityRole
}

func NewHandlerRole(role entitys.EntityRole) *handlerRole {
	return &handlerRole{role: role}
}

func (h *handlerRole) HandlerInsert(ctx *gin.Context) {
	var body schemas.SchemaRole
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorRole(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.role.EntityInsert(&body)

	if error.Type == "error_register_02" {
		helpers.APIResponse(ctx, "Insert new role failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Insert new role account success", http.StatusOK, nil)
}

func ValidatorRole(ctx *gin.Context, input schemas.SchemaRole, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	schema = gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Role_name",
				Message: "Role_name is required on body",
			},
		},
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
