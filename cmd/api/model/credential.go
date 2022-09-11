package model

import "github.com/google/uuid"

type Credential struct {
	ID             string
	Value          string
	LatestPassword string
}

func NewCredential(password string) Credential {
	return Credential{
		ID:             uuid.NewString(),
		Value:          password,
		LatestPassword: "123",
	}
}
