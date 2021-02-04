package main

import (
	"net/http"
	"time"

	_withdrawHttpHandler "github.com/cpartogi/withdrawdeposit/module/withdraw/handler/http"
	_withdrawRepo "github.com/cpartogi/withdrawdeposit/module/withdraw/store"
	_withdraw "github.com/cpartogi/withdrawdeposit/module/withdraw/usecase"

	_ "github.com/cpartogi/withdrawdeposit/docs"
	appInit "github.com/cpartogi/withdrawdeposit/init"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
	log "go.uber.org/zap"
)

func init() {
	// Start pre-requisite app dependencies
	appInit.StartAppInit()
}

func main() {

	// Get PG Conn Instance
	//	pgDb, err := appInit.ConnectToPGServer()
	pgDb, err := appInit.ConnectToMySqlServer()
	if err != nil {
		log.S().Fatal(err)
	}

	// init router
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is healthy")
	})

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	// DI: Repository & Usecase
	withdrawRepo := _withdrawRepo.NewStore(pgDb.DB)

	withdrawUc := _withdraw.NewWithdrawUsecase(withdrawRepo, timeoutContext)

	// End of DI Stepss

	_withdrawHttpHandler.NewWithdrawHandler(e, withdrawUc)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// start serve
	e.Logger.Fatal(e.Start(viper.GetString("api.port")))
}
