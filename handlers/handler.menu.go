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

type handlerMenu struct {
	menu entitys.EntityMenu
}

func NewHandlerMenu(menu entitys.EntityMenu) *handlerMenu {
	return &handlerMenu{menu: menu}
}

func (h *handlerMenu) HandlerInsert(ctx *gin.Context) {
	var body schemas.SchemaMenu
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorMenu(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.menu.EntityInsert(&body)

	if error.Type == "error_register_02" {
		helpers.APIResponse(ctx, "Insert new menu failed", error.Code, nil)

		return
	}

	helpers.APIResponse(ctx, "Insert new menu success", http.StatusOK, nil)
}

func ValidatorMenu(ctx *gin.Context, input schemas.SchemaMenu, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	schema = gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Menu_nama",
				Message: "Nama Menu is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Menu_link",
				Message: "Link Menu is required on body",
			},
		},
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
