package utils

import (
	"os"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// GetEnv returns app envorinment : e.g. development, production, staging, testing, etc
func GetEnv() string {
	return os.Getenv("APP_ENV")
}

// IsProductionEnv returns whether the app is running using production env
func IsProductionEnv() bool {
	return os.Getenv("APP_ENV") == "production"
}

func SetLowerAndAddSpace(str string) string {
	lower := matchFirstCap.ReplaceAllString(str, "${1} ${2}")
	lower = matchAllCap.ReplaceAllString(lower, "${1} ${2}")
	return strings.ToLower(lower)
}
