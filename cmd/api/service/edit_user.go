package service

import (
	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/model/request"
	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/model/storage"
	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/repository"

	"gorm.io/gorm"
)

func EditUser(user *request.RequestUserEdit, db *gorm.DB) (*storage.User, error) {
	return repository.EditUser(user, db)
}
