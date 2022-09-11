package router

import (
	"ambrosia-zeus-api/cmd/api/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func SetupEndpoints(db *gorm.DB) {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	router.POST("/users/create", handler.UserHandler{DB: db}.CreateUser)
	router.POST("/users/login", handler.UserHandler{DB: db}.LoginUser)
	router.PUT("/users/edit", handler.UserHandler{DB: db}.EditUser)

	err := router.Run()
	if err != nil {
		return
	}
}
