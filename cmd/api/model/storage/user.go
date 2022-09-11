package storage

import (
	"ambrosia-zeus-api/cmd/api/model/request"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID            string     `json:"id" gorm:"primaryKey"`
	UserType      string     `json:"user_type" gorm:"primaryKey" binding:"required"`
	FirstName     string     `json:"first_name" binding:"required"`
	LastName      string     `json:"last_name" binding:"required"`
	Username      string     `json:"username" binding:"required"`
	DocType       string     `json:"doc_type" binding:"required"`
	DocNumber     string     `json:"doc_number" binding:"required"`
	Phone         string     `json:"phone" binding:"required"`
	PersonalEmail string     `json:"personal_email" binding:"required"`
	DoB           time.Time  `json:"date_of_birth" binding:"required"`
	CredentialID  string     `json:"-" gorm:"size:70"`
	Credential    Credential `json:"-" binding:"required" gorm:"foreignKey:CredentialID;references:ID"`
}

func NewUserFromStruct(user *request.RequestUser) User {
	dateString := "02/01/2006"
	parsedDate, err := time.Parse(dateString, user.DoB)
	if err != nil {
		return User{}
	}

	newUser := User{
		ID:            uuid.NewString(),
		UserType:      user.UserType,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Username:      user.Username,
		DocType:       user.DocType,
		DocNumber:     user.DocNumber,
		Phone:         user.Phone,
		PersonalEmail: user.PersonalEmail,
		DoB:           parsedDate,
		Credential:    NewCredential(user.Credential),
	}
	return newUser
}
