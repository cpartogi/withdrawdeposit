package init

import (
	"github.com/spf13/viper"
	log "go.uber.org/zap"
)

// setupAuthHelper inits
func setupAuthHelper() {
	if !viper.IsSet("auth.private_key") || viper.GetString("auth.private_key") == "" {
		log.S().Fatal("auth.private_key can not be empty for better security on auth")
	}
}
