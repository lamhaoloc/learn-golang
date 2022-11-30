package common

import "net/http"

type AppError struct {
	StatusCode int   `json:"status_code"`
	ErrorStack error `json:"-"`
}

func InternalError(err error) *AppError {
	return &AppError{StatusCode: http.StatusInternalServerError, ErrorStack: err}
}
