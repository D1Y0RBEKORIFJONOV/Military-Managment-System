package v1

import (
	"context"
	"net/http"
	"strconv"

	"api_service/api/handlers/models"

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

	createdSoldier, err := h.serviceManager.SoldierService().CreateSoldiers(context.Background(), soldier)
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
// @Param        id path string true "soldier_id"
// @Success      200  {object}  models.Soldier
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handlerV1) GetSoldierByID(c *gin.Context) {
	id := c.Param("id")

	soldier, err := h.service.Soldier().GetByID(context.Background(), id)
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
// @Param        search query string false "search"
// @Success      200  {object}  models.SoldiersResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handlerV1) GetSoldiersList(c *gin.Context) {
	var (
		page, limit int
		search      string
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, h.log, "parsing page data", http.StatusInternalServerError, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "50")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, h.log, "parsing limit data", http.StatusInternalServerError, err.Error())
		return
	}

	search = c.Query("search")

	response, err := h.service.Soldier().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
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
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handlerV1) UpdateSoldier(c *gin.Context) {
	var (
		err           error
		updateSoldier = models.UpdateSoldier{}
	)

	id := c.Param("id")

	if err := c.ShouldBindJSON(&updateSoldier); err != nil {
		handleResponse(c, h.log, "error while decoding soldier data", http.StatusBadRequest, err.Error())
		return
	}

	if err = h.service.Soldier().Update(context.Background(), models.UpdateSoldier{
		ID:        id,
		Name:      updateSoldier.Name,
		Rank:      updateSoldier.Rank,
		Unit:      updateSoldier.Unit,
		CreatedAt: "",
		UpdatedAt: "",
	}); err != nil {
		handleResponse(c, h.log, "error while updating soldier data", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, "soldier data successfully updated")
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

	if err = h.service.Soldier().Delete(context.Background(), id); err != nil {
		handleResponse(c, h.log, "error while deleting soldier", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, "soldier successfully deleted")
}
