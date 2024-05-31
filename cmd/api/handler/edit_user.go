package handler

import (
	"fmt"
	"net/http"

	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/model/request"
	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/service"

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
