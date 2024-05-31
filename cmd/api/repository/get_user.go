package repository

import (
	"fmt"

	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/model/general"
	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/model/storage"

	"gorm.io/gorm"
)

func GetUser(userID string, db *gorm.DB) (*storage.User, error) {
	var retrievedUser *storage.User

	queryResult := db.First(&retrievedUser, "ID = ?", userID)
	if queryResult.Error != nil || queryResult.RowsAffected == 0 {
		return nil, general.UserNotFound
	}

	return retrievedUser, nil
}

func GetUserByUsername(username string, db *gorm.DB) (*storage.User, error) {
	var retrievedUser *storage.User

	queryResult := db.First(&retrievedUser, "Username = ?", username)
	if queryResult.Error != nil || queryResult.RowsAffected == 0 {
		fmt.Println(queryResult.Error)
		return nil, general.UserNotFound
	}

	return retrievedUser, nil
}
