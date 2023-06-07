package routers

import (
	"github.com/aminazadbakht1/golang-clean-web-api/api/handlers"
	"github.com/aminazadbakht1/golang-clean-web-api/config"
	"github.com/gin-gonic/gin"
)

func Country(r *gin.RouterGroup, cfg *config.Config){
	h:= handlers.NewCountryHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
}