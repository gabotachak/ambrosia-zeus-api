package handler

import (
	"ambrosia-zeus-api/cmd/api/model/request"
	"ambrosia-zeus-api/cmd/api/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler UserHandler) CreateUser(ctx *gin.Context) {
	var reqUser *request.RequestUser
	err := ctx.BindJSON(&reqUser)
	if err != nil {
		fmt.Println(err)
	}

	createdUser, serviceErr := service.CreateUser(reqUser, handler.DB)
	if serviceErr != nil {
	}

	ctx.JSON(http.StatusCreated, createdUser)
}
