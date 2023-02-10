// Package ohlc provides the ohlc service.
package ohlc

import (
	"context"

	"github.com/teezzan/ohlc/internal/controller/ohlc/data"
)

// Service defines the ohlc service.
type Service interface {
	CreateDataPoints(ctx context.Context, dataPoints [][]string) error
	GetDataPoints(ctx context.Context, payload data.GetOHLCRequest) ([]data.OHLCEntity, *int, error)
}
