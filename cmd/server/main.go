package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teezzan/ohlc/internal/config"
	ohlc "github.com/teezzan/ohlc/internal/controller/ohlc"
	ohlcRepository "github.com/teezzan/ohlc/internal/controller/ohlc/repository"
	"github.com/teezzan/ohlc/internal/database"
	"github.com/teezzan/ohlc/internal/router"
	"go.uber.org/zap"
)

func main() {
	//init dependencies
	conf := config.Init()

	// Database
	db, err := database.New(
		conf.Database.Username,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Database,
	)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Repositories
	ohlcRepo := ohlcRepository.NewRepository(db.SQL)

	// Services
	ohlcService := ohlc.NewService(zap.NewNop(), ohlcRepo)

	// HTTP Handlers
	ohlcHandler := ohlc.NewHTTPHandler(zap.NewNop(), ohlcService)

	// Router
	r := router.New(
		dummyHandler,
		ohlcHandler,
	)

	err = r.SetupRouter(gin.Default())
	if err != nil {
		panic(err)
	}

	// Listen and Server in 0.0.0.0:8080
	r.Run(fmt.Sprintf(":%d", conf.Server.Port))
}

func dummyHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}
