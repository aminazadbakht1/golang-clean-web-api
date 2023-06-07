package routers

import (
	"github.com/aminazadbakht1/golang-clean-web-api/api/handlers"
	"github.com/aminazadbakht1/golang-clean-web-api/api/middlewares"
	"github.com/aminazadbakht1/golang-clean-web-api/config"
	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup, cfg *config.Config){
	h := handlers.NewUsersHandler(cfg)
	router.POST("/send-otp", middlewares.OtpLimiter(cfg), h.SendOtp)
	router.POST("/login-by-mobile", h.RegisterLoginByMobileNumber)
	router.POST("/login-by-username", h.LoginByUsername)
	router.POST("/register-by-username", h.RegisterByUsername)
}