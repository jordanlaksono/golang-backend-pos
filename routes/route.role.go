package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jordanlaksono/golang-backend-pos.git/handlers"
	"github.com/jordanlaksono/golang-backend-pos.git/repositorys"
	"github.com/jordanlaksono/golang-backend-pos.git/services"
	"gorm.io/gorm"
)

func NewRouteRole(db *gorm.DB, router *gin.Engine) {

	repositoryRole := repositorys.NewRepositoryRole(db)
	serviceRole := services.NewServiceRole(repositoryRole)
	handlerRole := handlers.NewHandlerRole(serviceRole)

	route := router.Group("/api/v1/role")

	route.POST("/store", handlerRole.HandlerInsert)
}
