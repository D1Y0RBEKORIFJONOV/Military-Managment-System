package v1

import (
	"api_service/pkg/logger"
	l "api_service/pkg/logger"
	"context"
	"net/http"
	"strings"
	"time"

	soldiers1 "github.com/D1Y0RBEKORIFJONOV/Milltary-Managment-System-protos/gen/go/soldiers"
	"github.com/gin-gonic/gin"
)

// Register ...
// @Summary Register
// @Description Register - Api for registering users
// @Tags Register
// @Accept json
// @Produce json
// @Param Register body models.RegisterModel true "createRegisterModel"
// @Success 200 {object} models.RegisterModel
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /v1/register/ [post]
func (h *handlerV1) Register(c *gin.Context) {
	var body soldiers1.CreateSoldiersReq
	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to bind json", l.Error(err))
		return
	}
	body.Email = strings.TrimSpace(body.Email)
	body.Email = strings.ToLower(body.Email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	_, err = h.serviceManager.SoldierService().CreateSoldiers(ctx, &body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to create soldier", l.Error(err))
		return
	}
	h.log.Info("", logger.Any("soldier", body))

	c.JSON(http.StatusOK, true)
}

// Verification ...
// @Summary Verification
// @Description Verification - Api for registering users
// @Tags Register
// @Accept json
// @Produce json
// @Param code query int true "Code"
// @Param email query string true "Email"
// @Success 200 {object} models.RegisterModel
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /v1/verification/ [post]
func (h *handlerV1) Verification(c *gin.Context) {
	codeRegis := c.Query("code")
	emailRegis := c.Query("email")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	regis, err := h.serviceManager.SoldierService().RegisterUser(ctx, &soldiers1.RegisterReq{
		Email:      emailRegis,
		SecredCode: codeRegis,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, "email or code error")
		return
	}

	c.JSON(http.StatusOK, regis)
}

// LogIn ...
// @Summary LogIn
// @Description LogIn - Api for registering users
// @Tags Register
// @Accept json
// @Produce json
// @Param password query string true "Password"
// @Param email query string true "Email"
// @Success 200 {object} models.RegisterResponse
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /v1/login/ [post]
func (h *handlerV1) LogIn(c *gin.Context) {
	password := c.Query("password")
	email := c.Query("email")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.SoldierService().Login(ctx, &soldiers1.LoginReq{
		Password: password,
		Email:    email})
	if err != nil {
		c.JSON(http.StatusBadRequest, "password error")
		return
	}

	c.JSON(http.StatusOK, response.Token)
}
