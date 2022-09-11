package service

import (
	"ambrosia-zeus-api/cmd/api/model/request"
	"ambrosia-zeus-api/cmd/api/repository"
	"gorm.io/gorm"
)

func LoginUser(user *request.RequestLogin, db *gorm.DB) bool {
	return repository.LoginUser(user, db)
}
