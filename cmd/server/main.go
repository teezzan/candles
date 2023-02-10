package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	awsS3 "github.com/teezzan/ohlc/internal/client/s3"
	"github.com/teezzan/ohlc/internal/client/sqs"
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

	//Clients
	s3Client, err := awsS3.NewClient(context.Background(), zap.NewNop(), conf.S3Config)
	if err != nil {
		panic(err)
	}
	sqsClient, err := sqs.NewClient(context.Background(), zap.NewNop(), conf.SQSConfig)
	if err != nil {
		panic(err)
	}

	// Repositories
	ohlcRepo := ohlcRepository.NewRepository(db.SQL)

	// Services
	ohlcService := ohlc.NewService(zap.NewNop(), ohlcRepo, s3Client, sqsClient, conf.OHLCConfig)

	// HTTP Handlers
	ohlcHTTPHandler := ohlc.NewHTTPHandler(zap.NewNop(), ohlcService)

	// Router
	r := router.New(
		healthCheckHandlerFunc,
		ohlcHTTPHandler,
	)

	err = r.SetupRouter(gin.Default())
	if err != nil {
		panic(err)
	}

	// c := cron.New()
	// c.AddFunc("@every 2m", func() {
	ohlcService.GetAndProcessSQSMessage(context.Background())
	//  })
	// c.Start()

	// Listen and Server in 0.0.0.0:8080
	r.Run(fmt.Sprintf(":%d", conf.Server.Port))
}

func healthCheckHandlerFunc(c *gin.Context) {
	c.String(http.StatusOK, "Ok")
}
