package exceptions

import "net/http"

// AppError represents a structured API error.
type AppError struct {
	Status  int    `json:"-"`
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

func HttpError(httpStatus int) *AppError {
	return &AppError{
		Status:  httpStatus,
		Code:    uint32(httpStatus),
		Message: http.StatusText(httpStatus),
	}
}

var (
	ErrNotFound       = &AppError{Status: 200, Code: 1000404, Message: "resource not found"}
	ErrUnauthorized   = &AppError{Status: 200, Code: 1000401, Message: "authentication required"}
	ErrBadRequest     = &AppError{Status: 200, Code: 1000400, Message: "invalid request"}
	ErrInternalServer = &AppError{Status: 500, Code: 1000500, Message: "Internal server error"}
)
