package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/teezzan/candles/internal/controller/ohlc/data"
	E "github.com/teezzan/candles/internal/errors"
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

// InsertDataPoints inserts a slice of data.OHLCEntity rows into the ohlc_data table of the MySQL repository.
// It uses NamedExecContext to bind the values in the sql statement.
// It returns an error if it failed to insert the data into the table.
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

// GetDataPoints retrieves OHLC data points from the database for a given symbol and time range
// It returns a slice of OHLCEntity structs with data for a given symbol between the start and end times
// The result is paginated with page number and page size parameters
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

// GetProcessingStatus retrieves the processing status of a file from the database
// It returns a ProcessingStatusEntity struct with the status of the file
func (r *MySQLRepository) GetProcessingStatus(ctx context.Context, fileName string) (*data.ProcessingStatusEntity, error) {
	stmt := `
	SELECT
		file_name,
		status,
		error,
		created_at,
		updated_at
	FROM
		process_status
	WHERE
		file_name = ?
	`
	var status data.ProcessingStatusEntity

	err := r.GetContext(ctx, &status, stmt, fileName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, E.NewErrEntityNotFound("file", fileName)
		}
		return nil, err
	}
	return &status, nil
}

// InsertProcessingStatus inserts a ProcessingStatusEntity struct into the process_status table of the MySQL repository.
// It uses NamedExecContext to bind the values in the sql statement.
// It returns an error if it failed to insert the data into the table.
func (r *MySQLRepository) InsertProcessingStatus(ctx context.Context, status data.ProcessingStatusEntity) error {
	stmt := `
	INSERT INTO process_status
		(
			file_name,
			status
		) VALUES (
			:file_name,
			:status
		);
	`
	_, err := r.NamedExecContext(ctx, stmt, status)
	if err != nil {
		return err
	}
	return nil
}

// UpdateProcessingStatus updates the status of a file in the process_status table of the MySQL repository.
// It uses NamedExecContext to bind the values in the sql statement.
// It returns an error if it failed to update the status of the file in the table.
func (r *MySQLRepository) UpdateProcessingStatus(ctx context.Context, status data.ProcessingStatusEntity) error {
	stmt := `
	UPDATE process_status
	SET
		status = :status,
		error = :error,
	WHERE
		file_name = :file_name
	`
	_, err := r.NamedExecContext(ctx, stmt, status)
	if err != nil {
		return err
	}
	return nil
}

// RemoveStaleProcessingStatus removes stale entries from the process_status table of the MySQL repository.
// It uses NamedExecContext to bind the values in the sql statement.
// It returns an error if it failed to remove the stale entries from the table.
func (r *MySQLRepository) RemoveStaleProcessingStatus(ctx context.Context, staleTime time.Time) error {
	stmt := `
	DELETE FROM process_status
	WHERE
		updated_at <= ?
	`
	_, err := r.ExecContext(ctx, stmt, staleTime)
	if err != nil {
		return err
	}
	return nil
}
