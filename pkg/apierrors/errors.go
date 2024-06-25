package apierrors

import "net/http"

type ApiError struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func newApiError(status int, code string, message string, details string) *ApiError {
	return &ApiError{
		Status:  status,
		Code:    code,
		Message: message,
		Details: details,
	}
}

func InternalServerError(details string) *ApiError {
	return newApiError(
		http.StatusInternalServerError,
		"internal_server_error",
		"Internal Server Error",
		details)
}

func BadRequest(message, code, details string) *ApiError {
	return newApiError(
		http.StatusBadRequest,
		code,
		message,
		details)
}

func NotFound(message, code, details string) *ApiError {
	return newApiError(
		http.StatusNotFound,
		code,
		message,
		details)
}

func Unauthorized(message, code, details string) *ApiError {
	return newApiError(
		http.StatusUnauthorized,
		code,
		message,
		details)
}

func Forbidden(message, code, details string) *ApiError {
	return newApiError(
		http.StatusForbidden,
		code,
		message,
		details)
}
