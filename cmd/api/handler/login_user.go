package handler

import (
	"ambrosia-zeus-api/cmd/api/model/request"
	"ambrosia-zeus-api/cmd/api/model/response"
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

	loginErr := service.LoginUser(reqUser, handler.DB)

	if loginErr != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized login",
		})
	} else {
		loggedUser, userErr := service.GetUserByUsername(reqUser.Username, handler.DB)
		if userErr != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Error while fetching user",
			})
		}

		ctx.JSON(http.StatusOK, response.LoginResponse{
			UserID: loggedUser.ID,
		})
	}
}
