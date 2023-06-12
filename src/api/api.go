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
		// Test
		health := v1.Group("/health")
		test_router := v1.Group("/test" /*middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"})*/)

		// User
		users := v1.Group("/users")

		// Base
		countries := v1.Group("/countries", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		cities := v1.Group("/cities", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		files := v1.Group("/files", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		companies := v1.Group("/companies", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		colors := v1.Group("/colors", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		years := v1.Group("/years", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))

		// Property
		properties := v1.Group("/properties", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		propertyCategories := v1.Group("/property-categories", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))

		// Car
		carTypes := v1.Group("/car-types", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		gearboxes := v1.Group("/gearboxes", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModels := v1.Group("/car-models", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelColors := v1.Group("/car-model-colors", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelYears := v1.Group("/car-model-years", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelPriceHistories := v1.Group("/car-model-price-histories", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelImages := v1.Group("/car-model-images", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelProperties := v1.Group("/car-model-properties", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelComments := v1.Group("/car-model-comments", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin", "default"}))

		// Test
		routers.Health(health)
		routers.TestRouter(test_router)

		// User
		routers.User(users, cfg)

		// Base
		routers.Country(countries, cfg)
		routers.City(cities, cfg)
		routers.File(files, cfg)
		routers.Company(companies, cfg)
		routers.Color(colors, cfg)
		routers.Year(years, cfg)

		// Property
		routers.Property(properties, cfg)
		routers.PropertyCategory(propertyCategories, cfg)

		// Car
		routers.CarType(carTypes, cfg)
		routers.Gearbox(gearboxes, cfg)
		routers.CarModel(carModels, cfg)
		routers.CarModelColor(carModelColors, cfg)
		routers.CarModelYear(carModelYears, cfg)
		routers.CarModelPriceHistory(carModelPriceHistories, cfg)
		routers.CarModelImage(carModelImages, cfg)
		routers.CarModelProperty(carModelProperties, cfg)
		routers.CarModelComment(carModelComments, cfg)

		r.Static("/static", "./uploads")
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
