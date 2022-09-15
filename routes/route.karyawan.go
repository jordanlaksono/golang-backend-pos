package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jordanlaksono/golang-backend-pos.git/handlers"
	"github.com/jordanlaksono/golang-backend-pos.git/repositorys"
	"github.com/jordanlaksono/golang-backend-pos.git/services"
	"gorm.io/gorm"
)

func NewRouteKaryawan(db *gorm.DB, router *gin.Engine) {

	repositoryKaryawan := repositorys.NewRepositoryKaryawan(db)
	serviceKaryawan := services.NewServiceKaryawan(repositoryKaryawan)
	handlerKaryawan := handlers.NewHandlerKaryawan(serviceKaryawan)

	route := router.Group("/api/v1/karyawan")

	route.POST("/store", handlerKaryawan.HandlerInsert)
}
