package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupEndpoints() {
	router := gin.New()

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
	})

	err := router.Run()
	if err != nil {
		return
	}
}
