package helpers

import (
	"net/http"
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	case ErrInvalidCredentials:
		return http.StatusUnauthorized
	case ErrUserIsNotActive:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
