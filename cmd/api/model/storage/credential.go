package storage

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Credential struct {
	ID             string
	Value          string
	LatestPassword string
}

func NewCredential(credential string) Credential {
	hash := hashCredential(credential)

	return Credential{
		ID:             uuid.NewString(),
		Value:          hash,
		LatestPassword: "",
	}
}

func hashCredential(credential string) string {
	hashedCredential, err := bcrypt.GenerateFromPassword([]byte(credential), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	return string(hashedCredential)
}
