package apperrors

const (
	invalidItemData = -1
	recordNotFound  = -2
	invalidToken    = -3
	unauthorized    = -4
)

var (
	InvalidItemData = NewAppError(invalidItemData, "invalid item data")
	RecordNotFound  = NewAppError(recordNotFound, "record not found")
	InvalidToken    = NewAppError(invalidToken, "invalid token")
	Unauthorized    = NewAppError(unauthorized, "unauthorized")
)
