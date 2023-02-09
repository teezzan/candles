// Package ohlc provides the ohlc service.
package ohlc

import (
	"context"
)

// Service defines the ohlc service.
type Service interface {
	CreateOHLCPoints(ctx context.Context, dataPoints [][]string) error
}
