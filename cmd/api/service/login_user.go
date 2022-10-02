package service

import (
	"ambrosia-zeus-api/cmd/api/model/request"
	"ambrosia-zeus-api/cmd/api/repository"
	"gorm.io/gorm"
	"log"
)

func LoginUser(user *request.RequestLogin, db *gorm.DB) error {
	response := repository.LoginUser(user, db)
	if response != nil {
		log.Printf("Error on user login. Error: %s", response)
	}

	return response
}
