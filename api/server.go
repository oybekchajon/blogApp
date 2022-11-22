package api

import (
	"github.com/oybekchajon/blogApp/storage/postgres"
	"github.com/gin-gonic/gin"

	swagFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type handler struct {
	storage *postgres.DBManager
}



// @title           Swagger for book api
// @version         1.0
// @description     This is a book service api.
// @host      		localhost:8000
func NewServer(storage *postgres.DBManager) *gin.Engine {
	r := gin.Default()

	h := handler {
		storage: storage,
	}

	r.POST("/users",h.CreateUser)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swagFiles.Handler))
	return r
}