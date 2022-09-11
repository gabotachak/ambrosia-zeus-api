package service

import (
	"ambrosia-zeus-api/cmd/api/model/storage"
	"ambrosia-zeus-api/cmd/api/repository"
	"gorm.io/gorm"
)

func GetUser(userID string, db *gorm.DB) (*storage.User, error) {
	return repository.GetUser(userID, db)
}

func GetUserByUsername(userID string, db *gorm.DB) (*storage.User, error) {
	return repository.GetUserByUsername(userID, db)
}
