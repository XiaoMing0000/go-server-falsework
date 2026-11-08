package exceptions

// AppError represents a structured API error.
type AppError struct {
	Status  int    `json:"-"`
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

var (
	ErrNotFound     = &AppError{Status: 404, Code: 1000404, Message: "resource not found"}
	ErrUnauthorized = &AppError{Status: 401, Code: 1000401, Message: "authentication required"}
	ErrBadRequest   = &AppError{Status: 400, Code: 1000400, Message: "invalid request"}
)
