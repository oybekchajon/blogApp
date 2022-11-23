package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/oybekchajon/blogApp/storage/postgres"

	swagFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/oybekchajon/blogApp/api/docs" // for swagger
)

type handler struct {
	storage *postgres.DBManager
}

// @title           Swagger for user api
// @version         1.0
// @description     This is a user service api.
// @host      		localhost:8000
func NewServer(storage *postgres.DBManager) *gin.Engine {
	r := gin.Default()

	h := handler{
		storage: storage,
	}
	fmt.Println(h)

	r.GET("/users/:id", h.GetUser)
	r.POST("/users", h.CreateUser)
	r.PUT("/users/:id", h.UpdateUser)
	r.GET("users/", h.GetAllUsers)

	r.POST("/posts", h.CreatePost)
	r.GET("/posts/:id", h.GetPost)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swagFiles.Handler))
	return r
}
