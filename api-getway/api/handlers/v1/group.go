package v1

import (
	"context"
	"net/http"
	"strconv"

	group1 "github.com/D1Y0RBEKORIFJONOV/Milltary-Managment-System-protos/gen/go/group"
	"github.com/gin-gonic/gin"
)

// CreateGroup godoc
// @Router       /group [POST]
// @Summary      Creates a new group
// @Description  create a new group
// @Tags         group
// @Accept       json
// @Produce      json
// @Param        group body models.CreateGroupRequest true "group"
// @Success      201  {object}  models.Group
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handlerV1) CreateGroup(c *gin.Context) {
	var req group1.CreateGroupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handleResponse(c, h.log, "error while decoding group data", http.StatusBadRequest, err.Error())
		return
	}

	createdGroup, err := h.serviceManager.GroupService().CreateGroup(context.Background(), &req)
	if err != nil {
		handleResponse(c, h.log, "error while creating group", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "successfully created group", http.StatusCreated, createdGroup)
}

// GetGroupByID godoc
// @Router       /group/{id} [GET]
// @Summary      Get group by id
// @Description  Get group by id
// @Tags         group
// @Accept       json
// @Produce      json
// @Param        id path string true "group_id"
// @Success      200  {object}  models.Group
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handlerV1) GetGroupByID(c *gin.Context) {
	id := c.Param("id")

	group, err := h.serviceManager.GroupService().GetAllResourceTypes(context.Background(), &group1.GetAllServiceRequest{
		Field: "id",
		Value: id,
	})
	if err != nil {
		handleResponse(c, h.log, "error while get group by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "successfully", http.StatusOK, group)
}

// GetAllGroups godoc
// @Router       /groups [GET]
// @Summary      Get groups list
// @Description  Get groups list
// @Tags         group
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        field query string false "field"
// @Param        value query string false "value"
// @Param        sort_by query string false "sort_by"
// @Param        start_at query string false "start_at"
// @Param        end_at query string false "end_at"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handlerV1) GetAllGroups(c *gin.Context) {
	page, _ := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 64)
	limit, _ := strconv.ParseUint(c.DefaultQuery("limit", "100"), 10, 64)
	field := c.Query("field")
	value := c.Query("value")
	sort_by := c.Query("sort_by")
	start_at := c.Query("start_at")
	end_at := c.Query("end_at")

	response, err := h.serviceManager.GroupService().GetAllResourceTypes(context.Background(), &group1.GetAllServiceRequest{
		Field:   field,
		Value:   value,
		SordBy:  sort_by,
		StartAt: start_at,
		EndAt:   end_at,
		Page:    page,
		Limit:   limit,
	})
	if err != nil {
		handleResponse(c, h.log, "error while get groups list", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, response)
}

// UpdateGroup godoc
// @Router       /group/{id} [PUT]
// @Summary      Update group data
// @Description  Update group data
// @Tags         group
// @Accept       json
// @Produce      json
// @Param        id path string true "group_id"
// @Param        group body models.UpdateGroupRequest true "group"
// @Success      200  {object}  models.Group
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handlerV1) UpdateGroup(c *gin.Context) {
	id := c.Param("id")

	var req group1.UpdateGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleResponse(c, h.log, "error while decoding group data", http.StatusBadRequest, err.Error())
		return
	}

	req.GroupId = id
	group, err := h.serviceManager.GroupService().UpdateGroup(context.Background(), &req)
	if err != nil {
		handleResponse(c, h.log, "error while updating group data", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "group data successfully updated", http.StatusOK, group)
}

// DeleteGroup godoc
// @Router       /group/{id} [DELETE]
// @Summary      Delete group
// @Description  Delete group
// @Tags         group
// @Accept       json
// @Produce      json
// @Param        id path string true "group_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handlerV1) DeleteGroup(c *gin.Context) {
	id := c.Param("id")

	msg, err := h.serviceManager.GroupService().DeleteGroup(context.Background(), &group1.DeleteGroupRequest{
		GroupId: id,
	})
	if err != nil {
		handleResponse(c, h.log, "error while deleting group", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "group successfully deleted", http.StatusOK, msg)
}

// AddGroupSoldiers godoc
// @Router       /group/{id}/soldiers [POST]
// @Summary      Add soldiers to group
// @Description  Add soldiers to group
// @Tags         group
// @Accept       json
// @Produce      json
// @Param        id path string true "group_id"
// @Param        soldiers body models.AddGroupSoldersRequest true "soldiers"
// @Success      200  {object}  models.AddGroupSoldersResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handlerV1) AddGroupSoldiers(c *gin.Context) {
	id := c.Param("id")

	var req group1.AddGroupSoldersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleResponse(c, h.log, "error while decoding soldiers data", http.StatusBadRequest, err.Error())
		return
	}

	req.GroupId = id
	response, err := h.serviceManager.GroupService().AddGroupSolders(context.Background(), &req)
	if err != nil {
		handleResponse(c, h.log, "error while adding soldiers to group", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "soldiers successfully added to group", http.StatusOK, response)
}
