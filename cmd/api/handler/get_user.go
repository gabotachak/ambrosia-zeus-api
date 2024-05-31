package handler

import (
	"net/http"

	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/service"

	"github.com/gin-gonic/gin"
)

func (handler UserHandler) GetUser(ctx *gin.Context) {
	userID := ctx.Query("id")
	user, err := service.GetUser(userID, handler.DB)

	if err != nil {

	}

	ctx.JSON(http.StatusOK, user)
}

func (handler UserHandler) GetUserByUsername(ctx *gin.Context) {
	username := ctx.Query("username")
	user, err := service.GetUserByUsername(username, handler.DB)

	if err != nil {

	}

	ctx.JSON(http.StatusOK, user)
}
