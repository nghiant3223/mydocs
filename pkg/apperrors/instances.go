package apperrors

const (
	invalidItemData  = -1
	recordNotFound   = -2
	invalidToken     = -3
	unauthorized     = -4
	loginFailed      = -5
	invalidLoginData = -6
)

var (
	InvalidItemData  = NewAppError(invalidItemData, "invalid item data")
	RecordNotFound   = NewAppError(recordNotFound, "record not found")
	InvalidToken     = NewAppError(invalidToken, "invalid token")
	Unauthorized     = NewAppError(unauthorized, "unauthorized")
	LoginFailed      = NewAppError(loginFailed, "login failed")
	InvalidLoginData = NewAppError(invalidLoginData, "invalid login data")
)
