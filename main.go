package main

import (
	"net/http"

	appInit "github.com/cpartogi/withdrawdeposit/init"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func init() {
	//start app dependecies
	appInit.StartAppInit()
}

func main() {

	//db connection

	//init router
	e := echo.New()

	//middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is healthy")
	})

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	//	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	//depositUc := _auth.NewAuthUsecase(authRepo, timeoutContext)
	//	depositUc := _withdraw.NewDepositUsecase(authRepo, timeoutContext)

	// End of DI Stepss

	//	_withdrawHttpHandler.NewDepositHandler(e, depositUc)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// start serve
	e.Logger.Fatal(e.Start(viper.GetString("api.port")))
}
