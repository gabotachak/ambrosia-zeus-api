package repository

import (
	"fmt"

	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/model/request"
	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/model/storage"

	"gorm.io/gorm"
)

func EditUser(user *request.RequestUserEdit, db *gorm.DB) (*storage.User, error) {
	var retrievedUser *storage.User

	userResult := db.First(&retrievedUser, "username = ?", user.Username)
	if userResult.Error != nil || userResult.RowsAffected == 0 {
		return nil, nil
	}

	replacerValues := storage.NewUserFromEditStruct(user)
	updateResult := db.Model(&retrievedUser).Updates(replacerValues)
	if updateResult.Error != nil {
		fmt.Println(updateResult.Error)
		return nil, nil
	}

	return retrievedUser, nil
}
