package init

import (
	"strings"

	"github.com/cpartogi/withdrawdeposit/pkg/utils"
	"github.com/spf13/viper"
	log "go.uber.org/zap"
)

// setupMainConfig loads app config to viper
func setupMainConfig() {
	log.S().Info("Executing init/config")

	viper.SetConfigFile("config/app/development.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.S().Info("err: ", err)
	}

	if utils.IsProductionEnv() {
		viper.SetConfigFile("config/app/production.yml")
		err := viper.ReadInConfig()
		if err != nil {
			log.S().Info("err: ", err)
		}
	}

	if utils.IsFileorDirExist("main.yml") {
		log.S().Info("Local main.yml file is found, now assigning it with default config")
		viper.SetConfigFile("main.yml")
		err = viper.MergeInConfig()
		if err != nil {
			log.S().Info("err: ", err)
		}
	}

	viper.SetEnvPrefix(`app`)
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()

	log.S().Info("Config- APP_ENV: ", utils.GetEnv())
}
