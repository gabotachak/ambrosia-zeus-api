package repository

import (
	"ambrosia-zeus-api/cmd/api/model"
	"gorm.io/gorm"
)

func CreateUser(user *model.RequestUser, db *gorm.DB) (*model.User, error) {
	userToCreate := model.NewUserFromStruct(user)

	db.Create(&userToCreate)

	return &userToCreate, nil
}
