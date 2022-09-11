package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID            string     `json:"id" gorm:"primaryKey"`
	UserType      string     `json:"user_type" binding:"required"`
	FirstName     string     `json:"first_name" binding:"required"`
	LastName      string     `json:"last_name" binding:"required"`
	Username      string     `json:"username" binding:"required"`
	DocType       string     `json:"doc_type" binding:"required"`
	DocNumber     string     `json:"doc_number" binding:"required"`
	Phone         string     `json:"phone" binding:"required"`
	PersonalEmail string     `json:"personal_email" binding:"required"`
	DoB           time.Time  `json:"date_of_birth" binding:"required"`
	Credential    Credential `json:"credential" binding:"required" gorm:"foreignKey:UserCredential;references:ID"`
}

func NewUserFromStruct(user *RequestUser) *User {
	dateString := "02/01/2006"
	parsedDate, err := time.Parse(user.DoB, dateString)
	if err != nil {
		return nil
	}

	newUser := &User{
		ID:            uuid.NewString(),
		UserType:      user.UserType,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Username:      user.LastName,
		DocType:       user.DocType,
		DocNumber:     user.DocNumber,
		Phone:         user.Phone,
		PersonalEmail: user.PersonalEmail,
		DoB:           parsedDate,
		Credential:    NewCredential(user.Credential),
	}
	return newUser
}
