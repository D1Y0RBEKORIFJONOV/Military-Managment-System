package v1

// import (
// 	"api_service/api/handlers/models"
// 	compb "api_service/genproto/comment-service"
// 	postpb "api_service/genproto/post-service"
// 	pb "api_service/genproto/user-service"
// 	l "api_service/pkg/logger"
// 	"context"
// 	"github.com/gin-gonic/gin"
// 	"google.golang.org/protobuf/encoding/protojson"
// 	"net/http"
// 	"time"
// )

// // AdminCreatePost ...
// // @Summary AdminCreatePost
// // @Security ApiKeyAuth
// // @Description Api for creating a new post
// // @Tags admin
// // @Accept json
// // @Produce json
// // @Param Post body models.AdminPost true "Create post"
// // @Success 200 {object} models.AdminPost
// // @Failure 400 {object} models.StandardErrorModel
// // @Failure 500 {object} models.StandardErrorModel
// // @Router /v1/admin/create/post/ [post]
// func (h *handlerV1) AdminCreatePost(c *gin.Context) {
// 	var (
// 		body        models.AdminPost
// 		jsonMarshal protojson.MarshalOptions
// 	)
// 	jsonMarshal.UseProtoNames = true

// 	err := c.ShouldBindJSON(&body)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.log.Error("failed to bind json", l.Error(err))
// 		return
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
// 	defer cancel()

// 	response, err := h.serviceManager.PostService().Create(ctx, &postpb.Post{
// 		Title:    body.Title,
// 		ImageUrl: body.ImageUrl,
// 		Content:  body.Content,
// 		Id:       body.Id,
// 		OwnerId:  body.OwnerId,
// 		Category: body.Category,
// 	})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.log.Error("failed to create post", l.Error(err))
// 		return
// 	}
// 	c.JSON(http.StatusCreated, response)
// }

// // AdminUpdatePost ...
// // @Summary AdminUpdatePost
// // @Security ApiKeyAuth
// // @Description Api for creating a new post
// // @Tags admin
// // @Accept json
// // @Produce json
// // @Param Post body models.AdminPost true "update post"
// // @Param post_id query string true "id"
// // @Success 200 {object} models.AdminPost
// // @Failure 400 {object} models.StandardErrorModel
// // @Failure 500 {object} models.StandardErrorModel
// // @Router /v1/admin/update/post/ [put]
// func (h *handlerV1) AdminUpdatePost(c *gin.Context) {
// 	id := c.Query("id")
// 	var (
// 		body        models.AdminPost
// 		jsonMarshal protojson.MarshalOptions
// 	)
// 	jsonMarshal.UseProtoNames = true

// 	err := c.ShouldBindJSON(&body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.log.Error("failed to bind json", l.Error(err))
// 		return
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
// 	defer cancel()

// 	response, err := h.serviceManager.PostService().UpdatePost(ctx, &postpb.Post{
// 		Title:    body.Title,
// 		Category: body.Category,
// 		Content:  body.Content,
// 		ImageUrl: body.ImageUrl,
// 		Id:       id,
// 	})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.log.Error("failed to updated post", l.Error(err))
// 		return
// 	}
// 	c.JSON(http.StatusCreated, response)
// }

// // AdminUpdateUser updates user by id
// // @Summary AdminUpdateUser
// // @Security ApiKeyAuth
// // @Description Api for updating user
// // @Tags admin
// // @Accept json
// // @Produce json
// // @Param User body models.User true "createUserModel"
// // @Success 200 {object} models.User
// // @Failure 400 {object} models.StandardErrorModel
// // @Failure 500 {object} models.StandardErrorModel
// // @Router /v1/admin/update/user/ [put]
// func (h *handlerV1) AdminUpdateUser(c *gin.Context) {
// 	var (
// 		body        models.User
// 		jspbMarshal protojson.MarshalOptions
// 	)
// 	jspbMarshal.UseProtoNames = true

// 	err := c.ShouldBindJSON(&body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.log.Error("failed to bind json", l.Error(err))
// 		return
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
// 	defer cancel()

// 	_, refresh, err := h.jwthandler.GenerateAuthJWT()

// 	response, err := h.serviceManager.UserService().Update(ctx, &pb.User{
// 		FirstName:    body.FirstName,
// 		LastName:     body.LastName,
// 		Username:     body.UserName,
// 		Password:     body.Password,
// 		Email:        body.Email,
// 		RefreshToken: refresh,
// 	})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.log.Error("failed to update user", l.Error(err))
// 		return
// 	}

// 	c.JSON(http.StatusOK, response)
// }

// // AdminDeleteUser deletes user by id
// // @Summary AdminDeleteUser
// // @Security ApiKeyAuth
// // @Description Api for deleting user by id
// // @Tags admin
// // @Accept json
// // @Produce json
// // @Param id query string true "ID"
// // @Success 200 {object} models.User
// // @Failure 400 {object} models.StandardErrorModel
// // @Failure 500 {object} models.StandardErrorModel
// // @Router /v1/admin/delete/user [delete]
// func (h *handlerV1) AdminDeleteUser(c *gin.Context) {
// 	var jspbMarshal protojson.MarshalOptions
// 	jspbMarshal.UseProtoNames = true

// 	guid := c.Query("id")
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
// 	defer cancel()

// 	response, err := h.serviceManager.UserService().Delete(
// 		ctx, &pb.GetRequest{
// 			UserId: guid,
// 		})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.log.Error("failed to delete user", l.Error(err))
// 		return
// 	}

// 	c.JSON(http.StatusOK, response)
// }

// // AdminDeleteComment ...
// // @Summary AdminDeleteComment
// // @Security ApiKeyAuth
// // @Tags admin
// // @Accept json
// // @Produce json
// // @Param id query string true "ID"
// // @Success 200 {object} models.Comment
// // @Failure 400 {object} models.StandardErrorModel
// // @Failure 500 {object} models.StandardErrorModel
// // @Router /v1/admin/delete/comment [delete]
// func (h *handlerV1) AdminDeleteComment(c *gin.Context) {
// 	guid := c.Query("id")
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
// 	defer cancel()

// 	s, err := h.serviceManager.CommentService().DeleteComment(ctx, &compb.Get{
// 		Id: guid,
// 	})
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.log.Error("failed to query", l.Error(err))
// 		return
// 	}
// 	c.JSON(http.StatusOK, s)
// }

// // AdminGetUser
// // @Summary AdminGetUser
// // @Security ApiKeyAuth
// // @Description Api for getting user by id
// // @Tags user
// // @Accept json
// // @Produce json
// // @Param id query string true "ID"
// // @Success 200 {object} models.User
// // @Failure 400 {object} models.StandardErrorModel
// // @Failure 500 {object} models.StandardErrorModel
// // @Router /v1/admin/get/user [get]
// func (h *handlerV1) AdminGetUser(c *gin.Context) {
// 	id := c.Query("id")
// 	var jspbMarshal protojson.MarshalOptions
// 	jspbMarshal.UseProtoNames = true

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
// 	defer cancel()

// 	response, err := h.serviceManager.UserService().GetUser(
// 		ctx, &pb.GetRequest{
// 			UserId: id,
// 		})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.log.Error("failed to get user", l.Error(err))
// 		return
// 	}
// 	c.JSON(http.StatusOK, response)
// }

// // Create ...
// // @Summary Create
// // @Security ApiKeyAuth
// // @Description Api for creating a new user
// // @Tags admin
// // @Accept json
// // @Produce json
// // @Param User body models.User true "createUserModel"
// // @Success 200 {object} models.User
// // @Failure 400 {object} models.StandardErrorModel
// // @Failure 500 {object} models.StandardErrorModel
// // @Router /v1//admin/create/user [post]
// func (h *handlerV1) Create(c *gin.Context) {

// 	var (
// 		body        models.User
// 		jsonMarshal protojson.MarshalOptions
// 	)
// 	jsonMarshal.UseProtoNames = true

// 	err := c.ShouldBindJSON(&body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.log.Error("failed to bind json", l.Error(err))
// 		return
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
// 	defer cancel()

// 	response, err := h.serviceManager.UserService().Create(ctx, &pb.User{
// 		LastName:     body.FirstName,
// 		FirstName:    body.LastName,
// 		Username:     body.UserName,
// 		Password:     body.Password,
// 		Email:        body.Email,
// 		Id:           "",
// 		RefreshToken: "",
// 		Post:         nil,
// 	})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.log.Error("failed to create user", l.Error(err))
// 		return
// 	}

// 	c.JSON(http.StatusCreated, response)
// }
