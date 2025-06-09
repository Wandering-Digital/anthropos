package customerror

import (
	"errors"
	"net/http"
)

var (
	ErrWriteCache    = errors.New("failed to write cache")
	ErrReadCache     = errors.New("failed to read cache")
	ErrCacheNotFound = errors.New("not found in cache")

	ErrUserNotFound = errors.New("user not found")
	ErrUserConflict = errors.New("user conflict")
)

func NewCustomError(err error) *APIError {
	switch {
	case errors.Is(err, ErrUserNotFound):
		return NewError(http.StatusNotFound, "User not found", err)
	case errors.Is(err, ErrUserConflict):
		return NewError(http.StatusConflict, "User already exists", err)
	default:
		return NewError(http.StatusInternalServerError, "Something went wrong", err)
	}
}
