package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jordanlaksono/golang-backend-pos.git/entitys"
	"github.com/jordanlaksono/golang-backend-pos.git/helpers"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
	"net/http"
)

type handlerPenjualan struct {
	penjualan entitys.EntityPenjualan
}

func NewHandlerPenjualan(penjualan entitys.EntityPenjualan) *handlerPenjualan {
	return &handlerPenjualan{penjualan: penjualan}
}

func (h *handlerPenjualan) HandlerPenjualanByKode(ctx *gin.Context) {

	var body schemas.SchemaPenjualan

	res, error := h.penjualan.EntityPenjualanByKode(&body)

	if error.Type == "error_results_01" {
		helpers.APIResponse(ctx, "Pelanggan data not found", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Pelanggan data already to use", http.StatusOK, res)
}
