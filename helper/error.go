package helper

import "errors"

var (
	ErrWeekendNotAllowed  = errors.New("weekend attendance is not allowed")
	ErrInvalidTimestamp   = errors.New("invalid timestamp")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
