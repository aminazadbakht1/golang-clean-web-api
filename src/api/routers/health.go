package routers

import (
	"github.com/aminazadbakht1/golang-clean-web-api/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup){
	handler := handlers.NewHealthHandler()

	r.GET("/", handler.Health)
	r.POST("/", handler.HealthPost)
}