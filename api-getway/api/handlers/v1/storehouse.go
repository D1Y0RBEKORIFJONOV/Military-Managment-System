package v1

import (
	"api-test/api/handlers/models"
	jwt "api-test/api/tokens"
	"api-test/config"
	postpb "api-test/genproto/post-service"
	l "api-test/pkg/logger"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreatePost ...
// @Summary CreatePost
// @Security ApiKeyAuth
// @Description Api for creating a new post
// @Tags post
// @Accept json
// @Produce json
// @Param Post body models.Post true "create post"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/create/post/ [post]
func (h *handlerV1) CreatePost(c *gin.Context) {
	var (
		body        models.Post
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	tok := c.GetHeader("Authorization")
	claims, err := jwt.ExtractClaim(tok, []byte(config.Load().SignInKey))
	response, err := h.serviceManager.PostService().Create(ctx, &postpb.Post{
		Title:    body.Title,
		Category: body.Category,
		Content:  body.Content,
		ImageUrl: body.ImageUrl,
		OwnerId:  cast.ToString(claims["sub"]),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create post", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// GetPost ...
// @Summary GetPost
// @Security ApiKeyAuth
// @Description Api for get post
// @Tags post
// @Accept json
// @Produce json
// @Param id query string true "Id"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/get/post [get]
func (h *handlerV1) GetPost(c *gin.Context) {
	id := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	post, err := h.serviceManager.PostService().GetPost(ctx, &postpb.GetRequests{
		PostId: id,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to query", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, post)
}

// GetAllPosts gets user by id
// @Summary GetAllPosts
// @Security ApiKeyAuth
// @Description Api for getting posts
// @Tags post
// @Accept json
// @Produce json
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/posts/all [get]
func (h *handlerV1) GetAllPosts(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")

	reqPage, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("page error", l.Error(err))
		return
	}
	reqLimit, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("limit error", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resp, err := h.serviceManager.PostService().GetAllPost(ctx, &postpb.GetAllPostRequest{
		Page:  int64(reqPage),
		Limit: int64(reqLimit),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list users", l.Error(err))
	}
	c.JSON(http.StatusOK, resp)
}

// UpdatePost ...
// @Summary UpdatePost
// @Security ApiKeyAuth
// @Description Api for updating post
// @Tags post
// @Accept json
// @Produce json
// @Param Post body models.Post true "update post"
// @Param post_id query string true "id"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/update/post/ [put]
func (h *handlerV1) UpdatePost(c *gin.Context) {
	id := c.Query("id")
	var (
		body        models.Post
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	tok := c.GetHeader("Authorization")
	claims, err := jwt.ExtractClaim(tok, []byte(config.Load().SignInKey))

	posts, err := h.serviceManager.PostService().GetPostByOwnerId(ctx, &postpb.GetByOwnerIdRequest{
		OwnerId: cast.ToString(claims["sub"]),
	})
	n := 0
	for _, i := range posts.Posts {
		if i.Id != id {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			h.log.Error("request error", l.Error(err))
			return
		}
		n++
		break
	}
	if n == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad requests",
		})
		h.log.Error("request error", l.Error(err))
		return
	}

	response, err := h.serviceManager.PostService().UpdatePost(ctx, &postpb.Post{
		Title:    body.Title,
		Category: body.Category,
		Content:  body.Content,
		ImageUrl: body.ImageUrl,
		Id:       id,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to updated post", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// DeletePost gets user by id
// @Summary DeletePost
// @Security ApiKeyAuth
// @Description Api for deleting post
// @Tags post
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/delete/post [delete]
func (h *handlerV1) DeletePost(c *gin.Context) {
	id := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	tok := c.GetHeader("Authorization")
	claims, err := jwt.ExtractClaim(tok, []byte(config.Load().SignInKey))
	posts, err := h.serviceManager.PostService().GetPostByOwnerId(ctx, &postpb.GetByOwnerIdRequest{
		OwnerId: cast.ToString(claims["sub"]),
	})
	n := 0
	for _, i := range posts.Posts {
		if i.Id != id {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			h.log.Error("request error", l.Error(err))
			return
		}
		n++
		fmt.Println(i.Id, id)
		break
	}
	if n == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad requests",
		})
		h.log.Error("request error", l.Error(err))
		return
	}
	response, err := h.serviceManager.PostService().DeletePost(ctx, &postpb.GetRequests{PostId: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to deleted post", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

func (h *handlerV1) LikeDislike(c *gin.Context) {

}
