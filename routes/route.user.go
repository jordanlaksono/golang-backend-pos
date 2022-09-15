package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jordanlaksono/golang-backend-pos.git/handlers"
	"github.com/jordanlaksono/golang-backend-pos.git/repositorys"
	"github.com/jordanlaksono/golang-backend-pos.git/services"
	"gorm.io/gorm"
)

func NewRouteUser(db *gorm.DB, router *gin.Engine) {

	repositoryUser := repositorys.NewRepositoryUser(db)
	serviceUser := services.NewServiceUser(repositoryUser)
	handlerUser := handlers.NewHandlerUser(serviceUser)

	route := router.Group("/api/v1/auth")

	route.GET("/ping", handlerUser.HandlerPing)
	route.GET("/test", handlerUser.HandlerTest)
	route.GET("/results", handlerUser.HandlerResults)
	route.GET("/result/:id", handlerUser.HandlerResult)
	route.POST("/register", handlerUser.HandlerRegister)
	route.POST("/login", handlerUser.HandlerLogin)
}
