package handler

import (
	"ambrosia-zeus-api/cmd/api/model/request"
	"ambrosia-zeus-api/cmd/api/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler UserHandler) EditUser(ctx *gin.Context) {
	var reqUser *request.RequestUserEdit
	err := ctx.BindJSON(&reqUser)
	if err != nil {
		fmt.Println(err)
	}

	editedUser, serviceErr := service.EditUser(reqUser, handler.DB)
	if serviceErr != nil {

	}

	ctx.JSON(http.StatusCreated, editedUser)
}
