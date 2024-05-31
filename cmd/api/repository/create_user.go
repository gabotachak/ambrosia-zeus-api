package repository

import (
	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/model/request"
	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/model/storage"

	"gorm.io/gorm"
)

func CreateUser(user *request.RequestUser, db *gorm.DB) (*storage.User, error) {
	userToCreate := storage.NewUserFromCreateStruct(user)
	db.Create(&userToCreate)

	return &userToCreate, nil
}

func CheckUsername(username string, db *gorm.DB) bool {
	var retrievedUser *storage.User

	queryResult := db.First(&retrievedUser, "Username = ?", username)
	if queryResult.RowsAffected > 0 {
		return true
	}

	return false
}
