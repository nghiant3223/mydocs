package apperrors

type AppError struct {
	Code    int
	Message string
}

func NewAppError(code int, message string) error {
	return &AppError{Code: code, Message: message}
}

func (a *AppError) Error() string {
	return a.Message
}
