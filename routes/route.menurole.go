package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jordanlaksono/golang-backend-pos.git/handlers"
	"github.com/jordanlaksono/golang-backend-pos.git/repositorys"
	"github.com/jordanlaksono/golang-backend-pos.git/services"
	"gorm.io/gorm"
)

func NewRouteMenuRole(db *gorm.DB, router *gin.Engine) {

	repositoryMenuRole := repositorys.NewRepositoryMenuRole(db)
	serviceMenuRole := services.NewServiceMenuRole(repositoryMenuRole)
	handlerMenuRole := handlers.NewHandlerMenuRole(serviceMenuRole)

	route := router.Group("/api/v1/menurole")

	route.POST("/store", handlerMenuRole.HandlerInsert)
}
