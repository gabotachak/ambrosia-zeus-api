package handler

import (
	"ambrosia-zeus-api/cmd/api/model"
	"ambrosia-zeus-api/cmd/api/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type CreateUserHandler struct {
	DB *gorm.DB
}

func (handler CreateUserHandler) CreateUser(ctx *gin.Context) {
	var reqUser *model.RequestUser
	err := ctx.BindJSON(&reqUser)
	if err != nil {
		fmt.Println(err)
	}

	createdUser, err := service.CreateUser(reqUser, handler.DB)
	if err != nil {
	}

	ctx.JSON(http.StatusCreated, createdUser)
}
