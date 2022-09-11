package model

import (
	"crypto/sha256"
	"fmt"
	"github.com/google/uuid"
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
	hashFunc := sha256.New()
	hashFunc.Write([]byte(credential))

	hashedCredential := fmt.Sprintf("%x", hashFunc.Sum(nil))
	return hashedCredential
}
