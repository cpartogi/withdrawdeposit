package main

import (
	"net/http"
	"time"

	_authHttpHandler "github.com/cpartogi/withdrawdeposit/module/auth/handler/http"
	_authRepo "github.com/cpartogi/withdrawdeposit/module/auth/store"
	_auth "github.com/cpartogi/withdrawdeposit/module/auth/usecase"

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
	pgDb, err := appInit.ConnectToPGServer()
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
	authRepo := _authRepo.NewStore(pgDb.DB)

	authUc := _auth.NewAuthUsecase(authRepo, timeoutContext)

	// End of DI Stepss

	_authHttpHandler.NewAuthHandler(e, authUc)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// start serve
	e.Logger.Fatal(e.Start(viper.GetString("api.port")))
}
