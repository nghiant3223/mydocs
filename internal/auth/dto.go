package auth

type LoginRequestBody struct {
	Password string `json:"password"`
}

type LoginResponseBody struct {
	AccessToken string `json:"access_token"`
}
