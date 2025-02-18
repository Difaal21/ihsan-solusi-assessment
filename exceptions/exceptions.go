package exceptions

import "errors"

var (
	ErrNotFound            = errors.New("not found")
	ErrInternalServerError = errors.New("internal server error")
	ErrInvalidCredential   = errors.New("invalid credential")
	ErrConflict            = errors.New("conflict")
)
