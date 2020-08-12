package auth

import (
	"crypto/md5"
	"fmt"

	"github.com/nghiant3223/mydocs/pkg/apperrors"
	"github.com/nghiant3223/mydocs/pkg/tokenmanager"
)

type Service interface {
	Login(body LoginRequestBody) (LoginResponseBody, error)
}

type service struct {
	config       Config
	tokenManager tokenmanager.Manager
}

func NewService(config Config, manager tokenmanager.Manager) Service {
	return &service{config: config, tokenManager: manager}
}

func (s *service) Login(body LoginRequestBody) (LoginResponseBody, error) {
	digestPassword := hash(body.Password)
	if digestPassword != s.config.password {
		return LoginResponseBody{}, apperrors.LoginFailed
	}
	token, err := s.tokenManager.Generate(nil)
	if err != nil {
		return LoginResponseBody{}, err
	}
	var resBody LoginResponseBody
	resBody.AccessToken = token
	return resBody, nil
}

func hash(s string) string {
	b := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", b)
}
