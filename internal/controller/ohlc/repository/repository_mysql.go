package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/teezzan/ohlc/internal/controller/ohlc/data"
)

var _ Repository = (*MySQLRepository)(nil)

// MySQLRepository implements a MySQL repository.
type MySQLRepository struct {
	*sqlx.DB
}

// NewRepository initializes a new MySQL repository.
func NewRepository(db *sqlx.DB) *MySQLRepository {
	return &MySQLRepository{db}
}

// InsertDataPoints creates OHLC points
func (r *MySQLRepository) InsertDataPoints(ctx context.Context, rows []data.OHLCEntity) error {
	stmt := `
	INSERT INTO ohlc_data
		(
			time,
			symbol,
			open,
			high,
			low,
			close
		) VALUES (
			:time,
			:symbol,
			:open,
			:high,
			:low,
			:close
		);
	`
	_, err := r.NamedExecContext(ctx, stmt, rows)
	if err != nil {
		return err
	}
	return nil
}

// GetDataPoints returns OHLC points for a given symbol and time range with pagination
func (r *MySQLRepository) GetDataPoints(ctx context.Context, payload data.GetOHLCRequest) ([]data.OHLCEntity, error) {
	stmt := `
	SELECT
		time,
		symbol,
		open,
		high,
		low,
		close
	FROM
		ohlc_data
	WHERE
		symbol = ?
		AND time >= ?
		AND time <= ?
	ORDER BY time ASC
	LIMIT ?
	OFFSET ?
	`
	var ohlcPoints []data.OHLCEntity

	offset := (payload.PageNumber.Int64 - 1) * payload.PageSize.Int64
	startTime := time.Unix(payload.StartTime, 0)
	endTime := time.Unix(payload.EndTime.Int64, 0)

	err := r.SelectContext(ctx, &ohlcPoints, stmt, payload.Symbol, startTime, endTime, payload.PageSize, offset)
	if err != nil {
		return nil, err
	}
	return ohlcPoints, nil
}
