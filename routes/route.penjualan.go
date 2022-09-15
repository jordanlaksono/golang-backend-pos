package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jordanlaksono/golang-backend-pos.git/handlers"
	"github.com/jordanlaksono/golang-backend-pos.git/middlewares"
	"github.com/jordanlaksono/golang-backend-pos.git/repositorys"
	"github.com/jordanlaksono/golang-backend-pos.git/services"
	"gorm.io/gorm"
)

func NewRoutePenjualan(db *gorm.DB, router *gin.Engine) {
	repository := repositorys.NewRepositoryPenjualan(db)
	service := services.NewServicePenjualan(repository)
	handler := handlers.NewHandlerPenjualan(service)

	route := router.Group("/api/v1/penjualan")
	route.Use(middlewares.AuthToken())
	//route.Use(middlewares.AuthRole(map[string]bool{"Admin": true}))

	route.GET("/cekKodePenjualan", handler.HandlerPenjualanByKode)
}
