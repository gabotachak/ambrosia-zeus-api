package repository

import (
	"ambrosia-zeus-api/cmd/api/model/request"
	"ambrosia-zeus-api/cmd/api/model/storage"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginUser(user *request.RequestLogin, db *gorm.DB) bool {
	var retrievedUser *storage.User
	var retrievedCredential *storage.Credential

	userResult := db.First(&retrievedUser, "username = ?", user.Username)
	if userResult.Error != nil || userResult.RowsAffected == 0 {
		return false
	}

	credentialResult := db.First(&retrievedCredential, "ID = ?", retrievedUser.CredentialID)
	if credentialResult.Error != nil || credentialResult.RowsAffected == 0 {
		return false
	}

	passCheckErr := bcrypt.CompareHashAndPassword([]byte(retrievedCredential.Value), []byte(user.Password))
	if passCheckErr != nil {
		return false
	}

	return true
}
