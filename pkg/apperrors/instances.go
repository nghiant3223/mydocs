package apperrors

const (
	invalidItemData = -1
	recordNotFound  = -2
)

var (
	InvalidItemData = NewAppError(invalidItemData, "invalid item data")
	RecordNotFound  = NewAppError(recordNotFound, "record not found")
)
