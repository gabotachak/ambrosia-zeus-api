package handler

import (
	"ambrosia-zeus-api/cmd/api/model/request"
	"ambrosia-zeus-api/cmd/api/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler UserHandler) LoginUser(ctx *gin.Context) {
	var reqUser *request.RequestLogin
	err := ctx.BindJSON(&reqUser)
	if err != nil {
		fmt.Println(err)
	}

	response := service.LoginUser(reqUser, handler.DB)

	if response {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Successful login",
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized login",
		})
	}
}
