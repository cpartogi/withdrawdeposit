package middleware

import (
	"fmt"
	"strings"

	"github.com/cpartogi/withdrawdeposit/pkg/utils"
	"github.com/cpartogi/withdrawdeposit/schema/response"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// ApikeyMiddleware is
func ApikeyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		s := strings.Split(token, " ")

		if len(s) < 1 {
			utils.ErrorBadRequest(c, fmt.Errorf("header not found"), response.Default{})
		}
		if s[0] != "Apikey" {
			utils.ErrorBadRequest(c, fmt.Errorf("Invalid header"), response.Default{})
			return nil
		}
		if s[1] != viper.GetString("secret.api_key") {
			utils.ErrorBadRequest(c, fmt.Errorf("Invalid key"), response.Default{})
			return nil
		}

		return next(c)
	}
}
