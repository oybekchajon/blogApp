package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oybekchajon/blogApp/api/models"
	"github.com/oybekchajon/blogApp/storage/repo"
)

// @Security ApiKeyAuth
// @Router /posts [post]
// @Summary Create a post
// @Description Create a post
// @Tags post
// @Accept json
// @Produce json
// @Param post body models.CreatePostRequest true "post"
// @Success 201 {object} models.Post
// @Failure 500 {object} models.ResponseError
func (h *handler) CreatePost(c *gin.Context) {
	var (
		req models.CreatePostRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.CreatePost(&repo.Post{
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
		UserID:      req.UserID,
		CategoryID:  req.CategoryID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Router /posts/{id} [get]
// @Summary Get post by id
// @Description Get post by id
// @Tags post
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Post
// @Failure 500 {object} models.ResponseError
func (h *handler) GetPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.GetPost(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
