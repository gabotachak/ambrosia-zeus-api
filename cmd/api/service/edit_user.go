package service

import (
	"ambrosia-zeus-api/cmd/api/model/request"
	"ambrosia-zeus-api/cmd/api/model/storage"
	"ambrosia-zeus-api/cmd/api/repository"
	"gorm.io/gorm"
)

func EditUser(user *request.RequestUserEdit, db *gorm.DB) (*storage.User, error) {
	return repository.EditUser(user, db)
}
