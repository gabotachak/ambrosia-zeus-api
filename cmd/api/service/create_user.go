package service

import (
	"ambrosia-zeus-api/cmd/api/model"
	"ambrosia-zeus-api/cmd/api/repository"
	"gorm.io/gorm"
)

func CreateUser(user *model.RequestUser, db *gorm.DB) (*model.User, error) {
	createdUser, _ := repository.CreateUser(user, db)
	return createdUser, nil
}
