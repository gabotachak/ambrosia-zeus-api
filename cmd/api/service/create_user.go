package service

import (
	"ambrosia-zeus-api/cmd/api/model/request"
	"ambrosia-zeus-api/cmd/api/model/storage"
	"ambrosia-zeus-api/cmd/api/repository"
	"gorm.io/gorm"
)

func CreateUser(user *request.RequestUser, db *gorm.DB) (*storage.User, error) {
	createdUser, _ := repository.CreateUser(user, db)
	return createdUser, nil
}

func GenerateUsernameFromName(firstName string, lastName string) string {
	return ""
}
