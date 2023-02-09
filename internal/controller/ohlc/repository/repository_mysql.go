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
