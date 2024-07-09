package api

import (
	_ "api_service/api/docs"
	v1 "api_service/api/handlers/v1"
	token "api_service/api/tokens"
	"api_service/config"
	"api_service/pkg/logger"
	"api_service/service"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	Enforcer       casbin.Enforcer
	CasbinEnforcer *casbin.Enforcer
	ServiceManager service.IServiceManager
}

// @title welcome to
// @version 1.7
// @host localhost:9999

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()

	jwtHandler := token.JWTHandler{
		SignKey: option.Conf.SignInKey,
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
		Jwthandler:     jwtHandler,
	})

	// router.Use(casbinC.NewAuthorizer())
	api := router.Group("/v1")

	// Authorization
	api.POST("/verification", handlerV1.Verification)
	api.POST("/register", handlerV1.Register)
	api.POST("/login", handlerV1.LogIn)

	api.POST("/soldier", handlerV1.CreateSoldier)
	api.GET("/soldiers", handlerV1.GetAllSoldiers)
	api.GET("/soldier/{id}", handlerV1.GetSoldierByID)
	api.PUT("/soldier/{id} ", handlerV1.UpdateSoldier)
	api.DELETE("/soldier/{id}", handlerV1.DeleteSoldier)

	api.POST("/storehouse", handlerV1.CreateStorehouse)
	api.GET("/storehouses", handlerV1.GetAllStorehouses)
	api.GET("/storehouse/{id}", handlerV1.GetStorehouseByID)
	api.PUT("/storehouse/{id} ", handlerV1.UpdateStorehouse)
	api.DELETE("/storehouse/{id}", handlerV1.DeleteStorehouse)

	url := ginSwagger.URL("swagger/doc.json")

	api.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
