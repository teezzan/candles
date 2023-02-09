// Package ohlc provides the ohlc service.
package ohlc

import (
	"context"

	"github.com/teezzan/ohlc/internal/controller/ohlc/data"
)

// Service defines the ohlc service.
type Service interface {
	CreateOHLCPoints(ctx context.Context, dataPoints [][]string) error
	GetOHLCPoints(ctx context.Context, payload data.GetOHLCRequest) ([]data.OHLCEntity, error)
}
