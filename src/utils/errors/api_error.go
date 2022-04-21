package errors

import "net/http"

type ApiErr struct {
	Message string
	Status int
	Error string
}

func NewInternalServerError(message string) *ApiErr {
	return &ApiErr{
		Message: message,
		Status: http.StatusInternalServerError,
		Error: "internal_server_error",
	}
}

func NewBadRequestError(message string) *ApiErr {
	return &ApiErr{
		Message: message,
		Status: http.StatusBadRequest,
		Error: "bad_request",
	}
}
