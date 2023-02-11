// Package ohlc provides the ohlc service.
package ohlc

import (
	"context"

	"github.com/teezzan/ohlc/internal/controller/ohlc/data"
)

//go:generate moq -rm -out service_mock.go . Service

// Service defines the ohlc service.
type Service interface {
	CreateDataPoints(ctx context.Context, dataPoints [][]string) error
	GetDataPoints(ctx context.Context, payload data.GetOHLCRequest) ([]data.OHLCEntity, *int, error)
	GeneratePreSignedURL(ctx context.Context) (*data.GeneratePresignedURLResponse, error)
	GetAndProcessSQSMessage(ctx context.Context) error
	DownloadAndProcessCSV(ctx context.Context, filename string) error
}
