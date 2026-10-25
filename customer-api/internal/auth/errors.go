package auth

import "errors"

var (
	ErrUserDisabled = errors.New(
		"user account is disabled",
	)
)
