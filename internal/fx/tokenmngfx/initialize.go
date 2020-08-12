package tokenmngfx

import (
	"github.com/nghiant3223/mydocs/pkg/tokenmanager"
	"github.com/spf13/viper"
)

func provideTokenManager() tokenmanager.Manager {
	secret := viper.GetString("jwt.secret")
	lifetime := viper.GetDuration("jwt.lifetime")

	return tokenmanager.NewManager(secret, lifetime)
}
