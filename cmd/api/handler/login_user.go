package handler

import (
	"fmt"
	"net/http"

	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/model/request"
	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/model/response"
	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/service"

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
			UserID:      loggedUser.ID,
			Username:    loggedUser.Username,
			DisplayName: fmt.Sprintf("%s %s", loggedUser.FirstName, loggedUser.LastName),
		})
	}
}
