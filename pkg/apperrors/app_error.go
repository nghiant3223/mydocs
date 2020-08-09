package apperrors

type AppError struct {
	Status  int
	Message string
}

func NewAppError(httpCode int, message string) error {
	return &AppError{Status: httpCode, Message: message}
}

func (a *AppError) Error() string {
	return a.Message
}
