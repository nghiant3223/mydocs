package apperrors

import "net/http"

var (
	InvalidItemData = NewAppError(http.StatusBadRequest, "invalid item data")
)
