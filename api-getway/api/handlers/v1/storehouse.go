package v1

import (
	"context"
	"net/http"
	"strconv"

	storehouses1 "github.com/D1Y0RBEKORIFJONOV/Milltary-Managment-System-protos/gen/go/storehouses"
	"github.com/gin-gonic/gin"
)

// CreateStorehouse godoc
// @Router       /storehouse [POST]
// @Summary      Creates a new storehouse
// @Description  create a new storehouse
// @Tags         storehouse
// @Accept       json
// @Produce      json
// @Param        storehouse body models.CreateStorehouseReq true "storehouse"
// @Success      201  {object}  models.Storehouse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handlerV1) CreateStorehouse(c *gin.Context) {
	createStorehouse := storehouses1.CreateStorehouseReq{}

	if err := c.ShouldBindJSON(&createStorehouse); err != nil {
		handleResponse(c, h.log, "error while decoding storehouse data", http.StatusBadRequest, err.Error())
		return
	}

	createdStorehouse, err := h.serviceManager.StorehouseService().CreateStorehouse(context.Background(), &createStorehouse)
	if err != nil {
		handleResponse(c, h.log, "error while creating storehouse", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "successfully created storehouse", http.StatusCreated, createdStorehouse)
}

// GetStorehouseByID godoc
// @Router       /storehouse/{id} [GET]
// @Summary      Get storehouse by id
// @Description  Get storehouse by id
// @Tags         storehouse
// @Accept       json
// @Produce      json
// @Param        field query string false "field"
// @Param        value query string false "value"
// @Success      200  {object}  models.Storehouse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handlerV1) GetStorehouseByID(c *gin.Context) {
	fields := c.Query("fields")
	value := c.Query("value")

	storehouse, err := h.serviceManager.StorehouseService().GetStorehouse(context.Background(), &storehouses1.GetStorehouseReq{
		Fields: fields,
		Value:  value,
	})
	if err != nil {
		handleResponse(c, h.log, "error while get storehouse by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "successfully", http.StatusOK, storehouse)
}

// GetAllStorehouses godoc
// @Router       /storehouses [GET]
// @Summary      Get storehouses list
// @Description  Get storehouses list
// @Tags         storehouse
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        field query string false "field"
// @Param        value query string false "value"
// @Param        sort_by query string false "sort_by"
// @Param        started_at query string false "started_at"
// @Param        ended_at query string false "ended_at"
// @Success      200  {object}  models.GetAllStorehouseRes
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handlerV1) GetAllStorehouses(c *gin.Context) {
	page, _ := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	limit, _ := strconv.ParseInt(c.DefaultQuery("limit", "100"), 10, 64)
	field := c.Query("field")
	value := c.Query("value")
	sort_by := c.Query("sort_by")
	started_at := c.Query("started_at")
	ended_at := c.Query("ended_at")

	response, err := h.serviceManager.StorehouseService().GetAllStorehouse(context.Background(), &storehouses1.GetAllStorehouseReq{
		Filed:   field,
		Value:   value,
		Page:    page,
		Limit:   limit,
		SordBy:  sort_by,
		StartAt: started_at,
		EndAt:   ended_at,
	})

	if err != nil {
		handleResponse(c, h.log, "error while getting storehouses list", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, response)
}

// UpdateStorehouse godoc             dew
// @Router       /storehouse/{id} [PUT]
// @Summary      Update storehouse data
// @Description  Update storehouse data
// @Tags         storehouse
// @Accept       json
// @Produce      json
// @Param        id path string true "storehouse_id"
// @Param        storehouse body models.UpdateStorehouseReq true "storehouse"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handlerV1) UpdateStorehouse(c *gin.Context) {
	var updateStorehouse storehouses1.UpdateStorehouseReq

	id := c.Param("id")

	if err := c.ShouldBindJSON(&updateStorehouse); err != nil {
		handleResponse(c, h.log, "error while decoding storehouse data", http.StatusBadRequest, err.Error())
		return
	}

	updateStorehouse.Id = id

	updatedStorehouse, err := h.serviceManager.StorehouseService().UpdateStorehouse(context.Background(), &updateStorehouse)
	if err != nil {
		handleResponse(c, h.log, "error while updating storehouse data", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, updatedStorehouse)
}

// DeleteStorehouse godoc
// @Router       /storehouse/{id} [DELETE]
// @Summary      Delete storehouse
// @Description  Delete storehouse
// @Tags         storehouse
// @Accept       json
// @Produce      json
// @Param        id path string true "storehouse_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handlerV1) DeleteStorehouse(c *gin.Context) {
	id := c.Param("id")

	_, err := h.serviceManager.StorehouseService().DeleteStorehouse(context.Background(), &storehouses1.DeleteStorehouseReq{Id: id})
	if err != nil {
		handleResponse(c, h.log, "error while deleting storehouse", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, "storehouse successfully deleted")
}
