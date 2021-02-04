package init

import (
	"os"

	"github.com/cpartogi/withdrawdeposit/internal/db"
	"github.com/cpartogi/withdrawdeposit/pkg/utils"
	"github.com/spf13/viper"
	log "go.uber.org/zap"
)

// ConnectToPGServer is a function to init PostgreSQL connection
func ConnectToMySqlServer() (*db.MySqlDB, error) {
	if utils.IsProductionEnv() && (!viper.IsSet("database.mysql.password") || viper.GetString("database.mysql.password") == "") {
		log.S().Error("database.pg.password can not be empty!")
		os.Exit(1)
	}

	dbmysql, err := db.CreateMySqlConnection(map[string]string{
		"host":     viper.GetString(`database.mysql.host`),
		"port":     viper.GetString(`database.mysql.port`),
		"user":     viper.GetString(`database.mysql.user`),
		"password": viper.GetString(`database.mysql.password`),
		"dbname":   viper.GetString(`database.mysql.dbname`),
	})

	if err != nil {
		os.Exit(1)
	}

	return dbmysql, err
}
