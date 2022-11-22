package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oybekchajon/blogApp/api/models"
	"github.com/oybekchajon/blogApp/storage/postgres"
)

// @Router /user [post]
// @Summary Create a user
// @Description Create a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body UserRequest true "User"
// @Success 200 {object} postgres.User
// @Failure 500 {object} ResponseError
func (h *handler) CreateUser(ctx *gin.Context) {
	var u models.UserRequest
	err := ctx.ShouldBindJSON(&u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message" : err.Error(),
		})
		return
	}

	created, err := h.storage.CreateUser(&postgres.User{
		FirstName: "John",
		LastName: "Doe",
		PhoneNumber: "2412",
		Email: "JohnDoe@gmail.com",
		Gender: "male",
		Password: "123321",
		Username: "john1",
		ProfileImageUrl: "some_url",
		Type: "developer",
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message" : err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, created)
}