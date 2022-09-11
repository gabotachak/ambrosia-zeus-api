package repository

import (
	"ambrosia-zeus-api/cmd/api/model"
	"fmt"
	"gorm.io/gorm"
)

func CreateUser(user *model.RequestUser, db *gorm.DB) (*model.User, error) {
	result := db.Create(model.NewUserFromStruct(user))
	fmt.Println(result)

	return nil, nil
}
