package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jordanlaksono/golang-backend-pos.git/handlers"
	"github.com/jordanlaksono/golang-backend-pos.git/repositorys"
	"github.com/jordanlaksono/golang-backend-pos.git/services"
	"gorm.io/gorm"
)

func NewRouteMenu(db *gorm.DB, router *gin.Engine) {

	repositoryMenu := repositorys.NewRepositoryMenu(db)
	serviceMenu := services.NewServiceMenu(repositoryMenu)
	handlerMenu := handlers.NewHandlerMenu(serviceMenu)

	route := router.Group("/api/v1/menu")

	route.POST("/store", handlerMenu.HandlerInsert)
	route.GET("/results", handlerMenu.HandlerResults)
}
