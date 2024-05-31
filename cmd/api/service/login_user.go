package service

import (
	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/model/request"
	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/repository"

	"log"

	"gorm.io/gorm"
)

func LoginUser(user *request.RequestLogin, db *gorm.DB) error {
	response := repository.LoginUser(user, db)
	if response != nil {
		log.Printf("Error on user login. Error: %s", response)
	}

	return response
}
