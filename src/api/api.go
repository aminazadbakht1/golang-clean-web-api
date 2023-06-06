package api

import (
	"fmt"

	"github.com/aminazadbakht1/golang-clean-web-api/api/middlewares"
	"github.com/aminazadbakht1/golang-clean-web-api/api/routers"
	"github.com/aminazadbakht1/golang-clean-web-api/api/validation"
	"github.com/aminazadbakht1/golang-clean-web-api/config"
	"github.com/aminazadbakht1/golang-clean-web-api/docs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

func InitServer(cfg *config.Config) {
	r := gin.New()
	RegisterValidators()

	r.Use(middlewares.Cors(cfg))
	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest() /* middlewares.TestMiddleware() */)

	RegisterRoutes(r, cfg)
	RegisterSwagger(r, cfg)

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")
		users := v1.Group("/users")

		routers.Health(health)
		routers.TestRouter(test_router)
		routers.User(users, cfg)
	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		routers.Health(health)
	}
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
		val.RegisterValidation("password", validation.PasswordValidator, true)
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "Golang Web api"
	docs.SwaggerInfo.Description = "Golang Web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
