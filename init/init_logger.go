package init

import (
	"os"

	"go.uber.org/zap"
	log "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logPath = "./logs/go.log"
)

func setupLogger() {
	_, err := os.Stat("logs")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("logs", 0755)
		if errDir != nil {
			log.S().Fatal(err)
		}

	}

	os.OpenFile(logPath, os.O_RDONLY|os.O_CREATE, 0666)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.DisableStacktrace = true

	// using json format if app_env is production or development
	if os.Getenv("APP_ENV") == "production" || os.Getenv("APP_ENV") == "development" {
		config = zap.NewProductionConfig()
		config.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		config.DisableStacktrace = false
	}

	config.EncoderConfig.TimeKey = "timestamp"
	//	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderofLayout("2006-01-02 15:04:05.000000")
	config.OutputPaths = []string{"stdout", logPath}
	logger, _ := config.Build()
	zap.ReplaceGlobals(logger)
}
