package v1

import (
	"context"
	"net/http"
	"strconv"

	soldiers1 "github.com/D1Y0RBEKORIFJONOV/Milltary-Managment-System-protos/gen/go/soldiers"
	"github.com/gin-gonic/gin"
)

// CreateSoldier godoc
// @Router       /soldier [POST]
// @Summary      Creates a new soldier
// @Description  create a new soldier
// @Tags         soldier
// @Accept       json
// @Produce      json
// @Param        soldier body models.CreateSoldier true "soldier"
// @Success      201  {object}  models.Soldier
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handlerV1) CreateSoldier(c *gin.Context) {
	createSoldier := soldiers1.CreateSoldiersReq{}

	if err := c.ShouldBindJSON(&createSoldier); err != nil {
		handleResponse(c, h.log, "error while decoding soldier data", http.StatusBadRequest, err.Error())
		return
	}

	createdSoldier, err := h.serviceManager.SoldierService().CreateSoldiers(context.Background(), &createSoldier)
	if err != nil {
		handleResponse(c, h.log, "error while creating soldier", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "successfully created soldier", http.StatusCreated, createdSoldier)
}

// GetSoldierByID godoc
// @Router       /soldier/{id} [GET]
// @Summary      Get soldier by id
// @Description  Get soldier by id
// @Tags         soldier
// @Accept       json
// @Produce      json
// @Param        field query string false "field"
// @Param        value query string false "value"
// @Success      200  {object}  models.Soldier
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handlerV1) GetSoldierByID(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")

	soldier, err := h.serviceManager.SoldierService().GetSoldier(context.Background(), &soldiers1.GetSoldierReq{
		Filed: field,
		Value: value,
	})
	if err != nil {
		handleResponse(c, h.log, "error while get soldier by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "successfully", http.StatusOK, soldier)
}

// GetSoldiersList godoc
// @Router       /soldiers [GET]
// @Summary      Get soldiers list
// @Description  Get soldiers list
// @Tags         soldier
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        field query string false "field"
// @Param        value query string false "value"
// @Param        sort_by query string false "sort_by"
// @Param        started_at query string false "started_at"
// @Param        ended_at query string false "ended_at"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handlerV1) GetAllSoldiers(c *gin.Context) {
	page, _ := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	limit, _ := strconv.ParseInt(c.DefaultQuery("limit", "100"), 10, 64)
	field := c.Query("field")
	value := c.Query("value")
	sort_by := c.Query("sort_by")
	started_at := c.Query("started_at")
	ended_at := c.Query("ended_at")

	response, err := h.serviceManager.SoldierService().GetAllSoldiers(context.Background(), &soldiers1.GetAllSoldierReq{
		Filed:   field,
		Value:   value,
		SordBy:  sort_by,
		StartAt: started_at,
		EndAt:   ended_at,
		Page:    page,
		Limit:   limit,
	})
	if err != nil {
		handleResponse(c, h.log, "error while get soldiers list", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, response)
}

// UpdateSoldier godoc
// @Router       /soldier/{id} [PUT]
// @Summary      Update soldier data
// @Description  Update soldier data
// @Tags         soldier
// @Accept       json
// @Produce      json
// @Param        id path string true "soldier_id"
// @Param        soldier body models.UpdateSoldier true "soldier"
// @Success      200  {object}  models.Soldier
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handlerV1) UpdateSoldier(c *gin.Context) {
	id := c.Param("id")

	soldiers, err := h.serviceManager.SoldierService().GetSoldier(context.Background(), &soldiers1.GetSoldierReq{
		Filed: "id",
		Value: id,
	})
	if err != nil {
		handleResponse(c, h.log, "error while get soldier by id", http.StatusInternalServerError, err.Error())
		return
	}
	var updateSoldier = soldiers1.UpdateSoldierReq{
		FnName:    soldiers.FnName,
		LnName:    soldiers.LnName,
		Password:  soldiers.Password,
		BirthDay:  soldiers.BirhtDay,
		SoldersId: soldiers.Id,
	}
	if err := c.ShouldBindJSON(&updateSoldier); err != nil {
		handleResponse(c, h.log, "error while decoding soldier data", http.StatusBadRequest, err.Error())
		return
	}

	sldr, err := h.serviceManager.SoldierService().UpdateSoldier(context.Background(), &updateSoldier)
	if err != nil {
		handleResponse(c, h.log, "error while updating soldier data", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "soldier data successfully updated", http.StatusOK, sldr)
}

// DeleteSoldier godoc
// @Router       /soldier/{id} [DELETE]
// @Summary      Delete soldier
// @Description  Delete soldier
// @Tags         soldier
// @Accept       json
// @Produce      json
// @Param        id path string true "soldier_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handlerV1) DeleteSoldier(c *gin.Context) {
	var (
		err error
	)
	id := c.Param("id")
	ishard, _ := strconv.ParseBool((c.Query("ishard")))

	status, err := h.serviceManager.SoldierService().DeleteSoldier(context.Background(), &soldiers1.DeleteSoldierReq{
		SoldersId:    id,
		IsHardDelete: ishard,
	})
	if err != nil {
		handleResponse(c, h.log, "error while deleting soldier", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "soldier successfully deleted", http.StatusOK, status)
}
