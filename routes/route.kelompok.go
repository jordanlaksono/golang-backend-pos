package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jordanlaksono/golang-backend-pos.git/handlers"
	"github.com/jordanlaksono/golang-backend-pos.git/middlewares"
	"github.com/jordanlaksono/golang-backend-pos.git/repositorys"
	"github.com/jordanlaksono/golang-backend-pos.git/services"
	"gorm.io/gorm"
)

func NewRouteKelompok(db *gorm.DB, router *gin.Engine) {
	repository := repositorys.NewRepositoryKelompok(db)
	service := services.NewServiceKelompok(repository)
	handler := handlers.NewHandlerKelompok(service)

	route := router.Group("/api/v1/kelompok")
	route.Use(middlewares.AuthToken())
	//route.Use(middlewares.AuthRole(map[string]bool{"Admin": true}))

	route.POST("/create", handler.HandlerCreate)
	route.GET("/results", handler.HandlerResults)
	route.GET("/result/:id", handler.HandlerResult)
	route.DELETE("/delete/:id", handler.HandlerDelete)
	route.POST("/update", handler.HandlerUpdate)
}
