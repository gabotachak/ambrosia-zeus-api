package service

import (
	"strconv"
	"strings"

	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/model/request"
	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/model/storage"
	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/repository"

	"gorm.io/gorm"
)

func CreateUser(user *request.RequestUser, db *gorm.DB) (*storage.User, error) {
	user.Username = GenerateUsernameFromName(user.FirstName, user.LastName, db)

	createdUser, _ := repository.CreateUser(user, db)
	return createdUser, nil
}

func GenerateUsernameFromName(firstName string, lastName string, db *gorm.DB) string {
	finalFirstName := strings.ToLower(string([]rune{getRune(firstName, 0)})) //first char from name
	finalLastName := strings.ToLower(lastName)

	username := finalFirstName + finalLastName

	var usernameExists bool
	idx := 0

	for {
		usernameExists = repository.CheckUsername(username+strconv.Itoa(idx), db)
		if !usernameExists {
			break
		}
		idx++
	}

	return username + strconv.Itoa(idx)
}

func getRune(str string, idx int) rune {
	return []rune(str)[idx]
}
