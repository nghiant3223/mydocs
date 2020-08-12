package tokenmanager

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nghiant3223/mydocs/pkg/apperrors"
)

var (
	alg = jwt.SigningMethodHS256
)

const (
	claimExp string = "exp"
	claimIat string = "iat"
)

type Manager interface {
	Generate(map[string]interface{}) (string, error)
	Validate(string) (map[string]interface{}, error)
}

type manager struct {
	secret   string
	lifetime time.Duration
}

func NewManager(secret string, lifetime time.Duration) Manager {
	return &manager{secret: secret, lifetime: lifetime}
}

func (m *manager) Generate(args map[string]interface{}) (string, error) {
	claims := make(jwt.MapClaims)
	for k, v := range args {
		claims[k] = v
	}
	now := time.Now()
	claims[claimIat] = now
	claims[claimExp] = now.Add(m.lifetime).Unix()
	token := jwt.NewWithClaims(alg, claims)
	return token.SignedString([]byte(m.secret))
}

func (m *manager) Validate(tokenStr string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenStr, m.keyFunc)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, apperrors.InvalidToken
	}
	return claims, nil
}

func (m *manager) keyFunc(*jwt.Token) (interface{}, error) {
	return []byte(m.secret), nil
}
