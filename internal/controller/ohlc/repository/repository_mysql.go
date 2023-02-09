package repository

import (
	"context"

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

// GetohlcByUUID returns a ohlc by UUID.
func (r *MySQLRepository) GetohlcByUUID(ctx context.Context, uuid string) (*data.OHLCEntity, error) {

	return nil, nil
}

// CreateOHLCPoints creates OHLC points
func (r *MySQLRepository) CreateOHLCPoints(ctx context.Context, rows []data.OHLCEntity) error {
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

// GetOHLCPoints returns OHLC points for a given symbol and time range with pagination
func (r *MySQLRepository) GetOHLCPoints(ctx context.Context, payload data.GetOHLCRequest) ([]data.OHLCEntity, error) {
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
	LIMIT ?
	OFFSET ?
	`
	var ohlcPoints []data.OHLCEntity

	offset := (payload.PageNumber.Int64 - 1) * payload.PageSize.Int64
	err := r.SelectContext(ctx, &ohlcPoints, stmt, payload.Symbol, payload.StartTime, payload.EndTime, payload.PageSize, offset)
	if err != nil {
		return nil, err
	}
	return ohlcPoints, nil
}
