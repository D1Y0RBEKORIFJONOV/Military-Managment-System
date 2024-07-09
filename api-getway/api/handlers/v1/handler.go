package v1

import (
	"api-test/api/handlers/models"
	token "api-test/api/tokens"
	"api-test/config"
	"api-test/pkg/logger"
	"api-test/service"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	jwthandler     token.JWTHandler
	log            logger.Logger
	serviceManager service.IServiceManager
	cfg            config.Config
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Jwthandler     token.JWTHandler
	Logger         logger.Logger
	ServiceManager service.IServiceManager
	Cfg            config.Config
	Enforcer       casbin.Enforcer
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		jwthandler:     c.Jwthandler,
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
	}
}

func handleResponse(c *gin.Context, log logger.Logger, msg string, statusCode int, data interface{}) {
	resp := models.Response{}

	switch code := statusCode; {
	case code < 400:
		resp.Description = "OK"
		log.Info("~~~~> OK", logger.String("msg", msg), logger.Any("status", code))
	case code == 401:
		resp.Description = "Unauthorized"
	case code < 500:
		resp.Description = "Bad Request"
		log.Error("!!!!! BAD REQUEST", logger.String("msg", msg), logger.Any("status", code))
	default:
		resp.Description = "Internal Server Error"
		log.Error("!!!!! INTERNAL SERVER ERROR", logger.String("msg", msg), logger.Any("status", code))
	}

	resp.StatusCode = statusCode
	resp.Data = data

	c.JSON(resp.StatusCode, resp)
}
