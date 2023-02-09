package repository

import (
	"context"

	"github.com/teezzan/ohlc/internal/controller/ohlc/data"
)

// Repository defines the period repository.
type Repository interface {
	CreateOHLCPoints(ctx context.Context, rows []data.OHLCEntity) error
	GetOHLCPoints(ctx context.Context, payload data.GetOHLCRequest) ([]data.OHLCEntity, error)
}
