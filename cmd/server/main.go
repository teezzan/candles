package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"github.com/teezzan/ohlc/internal/client/s3"
	"github.com/teezzan/ohlc/internal/client/sqs"
	"github.com/teezzan/ohlc/internal/config"
	ohlc "github.com/teezzan/ohlc/internal/controller/ohlc"
	ohlcRepository "github.com/teezzan/ohlc/internal/controller/ohlc/repository"
	"github.com/teezzan/ohlc/internal/database"
	"github.com/teezzan/ohlc/internal/router"
	"go.uber.org/zap"
)

//	@title			Candles API
//	@description	This is API specification for Candels, a OHLC data API platform.
//	@version		1.0
//	@host			localhost:8080
//	@BasePath		/
func main() {
	//init dependencies
	conf := config.Init()

	logger, _ := zap.NewProduction()
	defer logger.Sync()
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

	//Clients
	s3Client, err := s3.NewClient(context.Background(), logger, conf.S3Config)
	if err != nil {
		panic(err)
	}
	sqsClient, err := sqs.NewClient(context.Background(), logger, conf.SQSConfig)
	if err != nil {
		panic(err)
	}

	// Repositories
	ohlcRepo := ohlcRepository.NewRepository(db.SQL)

	// Services
	ohlcService := ohlc.NewService(logger, ohlcRepo, s3Client, sqsClient, conf.OHLCConfig)

	// HTTP Handlers
	ohlcHTTPHandler := ohlc.NewHTTPHandler(logger, ohlcService)

	// Router
	r := router.New(
		healthCheckHandlerFunc,
		ohlcHTTPHandler,
	)

	err = r.SetupRouter(gin.Default())
	if err != nil {
		panic(err)
	}

	c := cron.New()
	c.AddFunc("@every 2m", func() {
		logger.Info("Processing SQS messages")
		ohlcService.GetAndProcessSQSMessage(context.Background())
	})
	c.Start()

	// Listen and Serve
	r.Run(fmt.Sprintf(":%d", conf.Server.Port))
}

func healthCheckHandlerFunc(c *gin.Context) {
	c.String(http.StatusOK, "Ok")
}
