// Package ohlc provides the ohlc service.
package ohlc

import (
	"context"

	"github.com/teezzan/candles/internal/controller/ohlc/data"
)

//go:generate moq -rm -out service_mock.go . Service

// Service defines the ohlc service.
type Service interface {
	CreateDataPoints(ctx context.Context, dataPoints [][]string) error
	GetDataPoints(ctx context.Context, payload data.GetOHLCRequest) ([]data.OHLCEntity, *int, error)
	GeneratePreSignedURL(ctx context.Context) (*data.GeneratePresignedURLResponse, error)
	GetAndProcessSQSMessage(ctx context.Context) error
	DownloadAndProcessCSV(ctx context.Context, filename string) error
	UpdateProcessingStatus(ctx context.Context, filename string, status data.ProcessingStatus, err error) error
	DeleteStaleProcessingStatus(ctx context.Context, days int) error
	GetProcessingStatus(ctx context.Context, filename string) (*data.ProcessingStatusEntity, error)
}
