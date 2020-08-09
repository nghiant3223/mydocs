package controller

type Response struct {
	ReturnCode int         `json:"return_code"`
	Message    *string     `json:"message"`
	Data       interface{} `json:"data"`
}
