package domain

import "github.com/labstack/echo/v4"

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(code int, msg string) *echo.HTTPError {
	return &echo.HTTPError{Code: code, Message: msg}
}
