package general

import "errors"

var (
	UserNotFound          = errors.New("user not found")
	CredentialNotFound    = errors.New("credential not found")
	InvalidHashComparison = errors.New("invalid hash comparison")
)
