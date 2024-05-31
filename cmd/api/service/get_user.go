package service

import (
	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/model/storage"
	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/repository"

	"gorm.io/gorm"
)

func GetUser(userID string, db *gorm.DB) (*storage.User, error) {
	return repository.GetUser(userID, db)
}

func GetUserByUsername(username string, db *gorm.DB) (*storage.User, error) {
	return repository.GetUserByUsername(username, db)
}
