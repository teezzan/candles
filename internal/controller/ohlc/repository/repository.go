package repository

import (
	"context"

	"github.com/teezzan/candles/internal/controller/ohlc/data"
)

//go:generate moq -rm -out repository_mock.go . Repository

// Repository defines the period repository.
type Repository interface {
	InsertDataPoints(ctx context.Context, rows []data.OHLCEntity) error
	GetDataPoints(ctx context.Context, payload data.GetOHLCRequest) ([]data.OHLCEntity, error)
}
