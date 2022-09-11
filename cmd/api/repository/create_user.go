package repository

import (
	"ambrosia-zeus-api/cmd/api/model/request"
	"ambrosia-zeus-api/cmd/api/model/storage"
	"gorm.io/gorm"
)

func CreateUser(user *request.RequestUser, db *gorm.DB) (*storage.User, error) {
	userToCreate := storage.NewUserFromStruct(user)

	db.Create(&userToCreate)

	return &userToCreate, nil
}
