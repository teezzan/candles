package repository

import (
	"context"
	"time"

	"github.com/teezzan/candles/internal/controller/ohlc/data"
)

//go:generate moq -rm -out repository_mock.go . Repository

// Repository defines the period repository.
type Repository interface {
	InsertDataPoints(ctx context.Context, rows []data.OHLCEntity) error
	GetDataPoints(ctx context.Context, payload data.GetOHLCRequest) ([]data.OHLCEntity, error)
	RemoveStaleProcessingStatus(ctx context.Context, staleTime time.Time) error
	UpdateProcessingStatus(ctx context.Context, status data.ProcessingStatusEntity) error
	InsertProcessingStatus(ctx context.Context, status data.ProcessingStatusEntity) error
	GetProcessingStatus(ctx context.Context, fileName string) (*data.ProcessingStatusEntity, error)
}
