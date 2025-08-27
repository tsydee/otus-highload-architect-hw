package apperrs

import "errors"

var (
	ErrConditionViolation = errors.New("condition violation")
	ErrUnauthorize        = errors.New("unauthorize")
	ErrAlreadyExist       = errors.New("already exist")
	ErrNotFound           = errors.New("not found")
)
