package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oybekchajon/blogApp/api/models"
	"github.com/oybekchajon/blogApp/storage/postgres"
)

// @Router /users [post]
// @Summary Create a user
// @Description Create a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.UserRequest true "User"
// @Success 201 {object} postgres.User
// @Failure 500 {object} models.ResponseError
func (h *handler) CreateUser(ctx *gin.Context) {
	var u models.UserRequest
	err := ctx.ShouldBindJSON(&u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	created, err := h.storage.CreateUser(&postgres.User{
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		PhoneNumber:     u.PhoneNumber,
		Email:           u.Email,
		Gender:          u.Gender,
		Password:        u.Password,
		Username:        u.Username,
		ProfileImageUrl: u.ProfileImageUrl,
		Type:            u.Type,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, created)
}

// @Router /users/{id} [get]
// @Summary Get user by id
// @Description Get user by id
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 201 {object} postgres.User
// @Failure 500 {object} models.ResponseError
func (h *handler) GetUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := h.storage.GetUser(int(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, user)
}




// @Router /users/{id} [put]
// @Summary Update user
// @Description Update user
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param user body models.UserRequest true "User"
// @Success 201 {object} postgres.User
// @Failure 500 {object} models.ResponseError
func (h *handler) UpdateUser(ctx *gin.Context) {
	var u models.UserRequest
	err := ctx.ShouldBindJSON(&u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	updated, err := h.storage.UpdateUser(&postgres.User{
		ID: id,
		FirstName: u.FirstName,
		LastName: u.LastName,
		PhoneNumber: u.Password,
		Email: u.Email,
		Gender: u.Gender,
		Password: u.Password,
		Username: u.Username,
		ProfileImageUrl: u.ProfileImageUrl,
		Type: u.Type,
	})

	if err !=nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message" : err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updated)
}

// @Router /users [get]
// @Summary Get all users
// @Description Get all users
// @Tags user
// @Accept json
// @Produce json
// @Param filter query models.GetAllParams false "Filter"
// @Success 200 {object} models.GetAllUsersResponse
// @Failure 500 {object} models.ResponseError
func(h *handler) GetAllUsers(ctx *gin.Context) {
	req, err := validateGetAllParams(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.storage.GetAll(&postgres.GetAllUsersParams{
		Page: req.Page,
		Limit: req.Limit,
		Search: req.Search,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, result)
}



func validateGetAllParams(c *gin.Context) (*models.GetAllParams, error) {
	var (
		limit int = 10
		page  int = 1
		err   error
	)

	if c.Query("limit") != "" {
		limit, err = strconv.Atoi(c.Query("limit"))
		if err != nil {
			return nil, err
		}
	}

	if c.Query("page") != "" {
		page, err = strconv.Atoi(c.Query("page"))
		if err != nil {
			return nil, err
		}
	}

	return &models.GetAllParams{
		Limit:  int32(limit),
		Page:   int32(page),
		Search: c.Query("search"),
	}, nil
}
