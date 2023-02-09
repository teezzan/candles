package ohlc

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/teezzan/ohlc/internal/config"
	"github.com/teezzan/ohlc/internal/controller/ohlc/data"
	"github.com/teezzan/ohlc/internal/controller/ohlc/repository"
	E "github.com/teezzan/ohlc/internal/errors"
	"github.com/teezzan/ohlc/internal/null"
	"go.uber.org/zap"
)

var _ Service = (*DefaultService)(nil)

type DefaultService struct {
	logger               *zap.Logger
	repository           repository.Repository
	discardInCompleteRow bool
	defaulDataPointLimit int
}

func NewService(
	logger *zap.Logger,
	repository repository.Repository,
	ohlcConf config.OHLCConfig,
) *DefaultService {
	return &DefaultService{
		logger:               logger,
		repository:           repository,
		discardInCompleteRow: ohlcConf.DiscardInCompleteRow,
		defaulDataPointLimit: ohlcConf.DefaultDataPointLimit,
	}
}

// CreateOHLCPoints creates OHLC points
func (s *DefaultService) CreateOHLCPoints(ctx context.Context, dataPoints [][]string) error {
	if len(dataPoints) == 0 {
		return nil
	}

	header := dataPoints[0]
	fieldIndexes := getFieldTitleIndex(header)
	if fieldIndexes.IsInComplete() {
		return E.NewErrInvalidArgument("Invalid CSV header")
	}

	ohlcPoints := make([]data.OHLCEntity, 0, len(dataPoints)-1)
	for _, row := range dataPoints[1:] {
		d, err := getOHLCPoint(row, fieldIndexes)
		if err != nil {
			if s.discardInCompleteRow {
				s.logger.Warn("Discarding incomplete row", zap.Error(err))
			} else {
				return err
			}
		}
		ohlcPoints = append(ohlcPoints, *d)
	}

	err := s.repository.CreateOHLCPoints(ctx, ohlcPoints)
	if err != nil {
		return err
	}
	return nil
}

func getFieldTitleIndex(header []string) data.OHLCFieldIndexes {
	v := data.DefaultOHLCFieldIndexes
	for i, field := range header {
		switch field {
		case data.OpenFieldName.String():
			d := i
			v.Open.Index = &d
		case data.HighFieldName.String():
			d := i
			v.High.Index = &d
		case data.LowFieldName.String():
			d := i
			v.Low.Index = &d
		case data.CloseFieldName.String():
			d := i
			v.Close.Index = &d
		case data.UnixFieldName.String():
			d := i
			v.Unix.Index = &d
		case data.SymbolFieldName.String():
			d := i
			v.Symbol.Index = &d
		}
	}
	return v
}

func getOHLCPoint(row []string, fieldIndexes data.OHLCFieldIndexes) (*data.OHLCEntity, error) {
	var d data.OHLCEntity

	if fieldIndexes.Symbol.Index != nil {
		t := row[*fieldIndexes.Symbol.Index]
		d.Symbol = strings.TrimSpace(t)
	}

	if fieldIndexes.Unix.Index != nil {
		t := row[*fieldIndexes.Unix.Index]
		i, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			return nil, err
		}
		d.Time = time.Unix(i, 0)
	}

	if fieldIndexes.Open.Index != nil {
		t := row[*fieldIndexes.Open.Index]
		val, err := strconv.ParseFloat(strings.TrimSpace(t), 64)
		if err != nil {
			return nil, err
		}
		d.Open = val
	}

	if fieldIndexes.High.Index != nil {
		t := row[*fieldIndexes.High.Index]
		val, err := strconv.ParseFloat(strings.TrimSpace(t), 64)
		if err != nil {
			return nil, err
		}
		d.High = val
	}

	if fieldIndexes.Low.Index != nil {
		t := row[*fieldIndexes.Low.Index]
		val, err := strconv.ParseFloat(strings.TrimSpace(t), 64)
		if err != nil {
			return nil, err
		}
		d.Low = val
	}

	if fieldIndexes.Close.Index != nil {
		t := row[*fieldIndexes.Close.Index]
		val, err := strconv.ParseFloat(strings.TrimSpace(t), 64)
		if err != nil {
			return nil, err
		}
		d.Close = val
	}

	if d.IsInComplete() {
		return nil, E.NewErrInvalidArgument("Invalid CSV row")
	}

	return &d, nil
}

// GetOHLCPoints returns OHLC points for a given symbol and time range
func (s *DefaultService) GetOHLCPoints(ctx context.Context, payload data.GetOHLCRequest) ([]data.OHLCEntity, error) {
	if payload.Symbol == "" {
		return nil, E.NewErrInvalidArgument("symbol is required")
	}
	if payload.StartTime.IsZero() {
		return nil, E.NewErrInvalidArgument("from is required")
	}
	if payload.EndTime.Valid && payload.EndTime.Time.Before(payload.StartTime) {
		return nil, E.NewErrInvalidArgument("to must be greater than from")
	}
	if payload.PageSize.Valid && payload.PageSize.Int64 <= 0 {
		return nil, E.NewErrInvalidArgument("page size must be greater than 0")
	}
	if payload.PageNumber.Valid && payload.PageNumber.Int64 <= 0 {
		return nil, E.NewErrInvalidArgument("page number must be greater than 0")
	}

	if !payload.PageSize.Valid {
		payload.PageSize = null.NewInt(s.defaulDataPointLimit)
	}
	if !payload.PageNumber.Valid {
		payload.PageNumber = null.NewInt(1)
	}
	if !payload.EndTime.Valid {
		payload.EndTime.Time = time.Now()
	}
	// calculate offset
	return s.repository.GetOHLCPoints(ctx, payload)
}
