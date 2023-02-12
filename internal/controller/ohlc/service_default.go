package ohlc

import (
	"context"
	"encoding/csv"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/teezzan/candles/internal/client/s3"
	"github.com/teezzan/candles/internal/client/sqs"
	"github.com/teezzan/candles/internal/config"
	"github.com/teezzan/candles/internal/controller/ohlc/data"
	"github.com/teezzan/candles/internal/controller/ohlc/repository"
	E "github.com/teezzan/candles/internal/errors"
	"github.com/teezzan/candles/internal/null"
	"github.com/teezzan/candles/internal/util"
	"go.uber.org/zap"
)

var _ Service = (*DefaultService)(nil)

type DefaultService struct {
	logger               *zap.Logger
	repository           repository.Repository
	s3Client             s3.Client
	sqsClient            sqs.Client
	discardInCompleteRow bool
	defaulDataPointLimit int
}

func NewService(
	logger *zap.Logger,
	repository repository.Repository,
	s3Client s3.Client,
	sqsClient sqs.Client,
	ohlcConf config.OHLCConfig,
) *DefaultService {
	return &DefaultService{
		logger:               logger,
		repository:           repository,
		discardInCompleteRow: ohlcConf.DiscardInCompleteRow,
		defaulDataPointLimit: ohlcConf.DefaultDataPointLimit,
		s3Client:             s3Client,
		sqsClient:            sqsClient,
	}
}

// CreateDataPoints creates OHLCEntities from a 2D array of strings and inserts them into the repository.
// The first row is expected to contain the header. If a row is incomplete, it can either be discarded
// or return an error based on the value of `discardInCompleteRow`.
func (s *DefaultService) CreateDataPoints(ctx context.Context, dataPoints [][]string) error {
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
		d, err := extractDataPoint(row, fieldIndexes)
		if err != nil {
			if s.discardInCompleteRow {
				s.logger.Warn("Discarding incomplete row", zap.Error(err))
			} else {
				return err
			}
		} else {
			ohlcPoints = append(ohlcPoints, *d)
		}
	}

	err := s.repository.InsertDataPoints(ctx, ohlcPoints)
	if err != nil {
		return err
	}
	return nil
}

// getFieldTitleIndex returns a `data.OHLCFieldIndexes` containing the index positions of OHLC and Unix fields in a given header
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

// extractDataPoint takes in a string slice representing a row from a CSV file and a data.OHLCFieldIndexes object,
// parses the values from the row and returns a pointer to a data.OHLCEntity object if successful,
// or an error if the row does not contain the necessary data.
func extractDataPoint(row []string, fieldIndexes data.OHLCFieldIndexes) (*data.OHLCEntity, error) {
	var d data.OHLCEntity

	if len(row) != reflect.TypeOf(fieldIndexes).NumField() {
		return nil, E.NewErrInvalidArgument("Invalid CSV row")
	}

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

// GetDataPoints returns a slice of OHLCEntity representing the requested open-high-low-close data points for a specific symbol.
// It validates the inputs such as symbol, start and end time, page size and page number and returns an error if they are not valid.
// The page size and page number are optional and default to defaultDataPointLimit and 1 respectively if not provided.
// The end time is also optional and defaults to the current time if not provided.
// The result is based on the data obtained from the repository.
func (s *DefaultService) GetDataPoints(ctx context.Context, payload data.GetOHLCRequest) ([]data.OHLCEntity, *int, error) {
	if payload.Symbol == "" {
		return nil, nil, E.NewErrInvalidArgument("symbol is required")
	}
	if payload.StartTime <= 0 {
		return nil, nil, E.NewErrInvalidArgument("from is required")
	}
	if payload.EndTime.Valid && payload.EndTime.Int64 < payload.StartTime {
		return nil, nil, E.NewErrInvalidArgument("to must be greater than from")
	}
	if payload.PageSize.Valid && payload.PageSize.Int64 <= 0 {
		return nil, nil, E.NewErrInvalidArgument("page size must be greater than 0")
	}
	if payload.PageNumber.Valid && payload.PageNumber.Int64 <= 0 {
		return nil, nil, E.NewErrInvalidArgument("page number must be greater than 0")
	}

	if !payload.PageSize.Valid {
		payload.PageSize = null.NewInt(s.defaulDataPointLimit)
	}
	if !payload.PageNumber.Valid {
		payload.PageNumber = null.NewInt(1)
	}
	if !payload.EndTime.Valid {
		payload.EndTime = null.NewInt64(time.Now().Unix())
	}

	data, err := s.repository.GetDataPoints(ctx, payload)
	if err != nil {
		return nil, nil, err
	}
	return data, payload.PageNumber.AsRef(), nil
}

// GeneratePreSignedURL generates a presigned URL for uploading a file to S3.
// It returns the generated URL and filename.
func (s *DefaultService) GeneratePreSignedURL(ctx context.Context) (*data.GeneratePresignedURLResponse, error) {
	filename := fmt.Sprintf("%s.csv", util.GenerateUUID())
	url, err := s.s3Client.GeneratePresignedURL(ctx, filename)
	if err != nil {
		return nil, err
	}

	return &data.GeneratePresignedURLResponse{
		URL:      url,
		Filename: filename,
	}, nil
}

// GetAndProcessSQSMessage retrieves filenames from SQS messages, logs them and creates Goroutines
// to process the files in parallel using the DownloadAndProcessCSV method.
// It returns an error if any occurred during the retrieval of filenames from SQS messages.
func (s *DefaultService) GetAndProcessSQSMessage(ctx context.Context) error {
	filenames, err := s.sqsClient.GetFilenamesFromMessages(ctx)
	s.logger.Info("filenames are", zap.Any("filenames", filenames))
	if err != nil {
		return err
	}
	if len(filenames) == 0 {
		return nil
	}
	for _, filename := range filenames {
		s.UpdateProcessingStatus(ctx, filename, data.ProcessingStatusInProgress, nil)
		// Create Goroutines to process the files in parallel
		go s.DownloadAndProcessCSV(ctx, filename)
	}
	return nil
}

// DownloadAndProcessCSV retrieves a large CSV object from S3 and processes the data to create data points.
// If an error occurs while downloading the object from S3 or processing the data, it will be returned.
func (s *DefaultService) DownloadAndProcessCSV(ctx context.Context, filename string) error {
	s3FileData, err := s.s3Client.DownloadLargeObject(ctx, filename)
	if err != nil {
		s.UpdateProcessingStatus(ctx, filename, data.ProcessingStatusFailed, err)
		return err
	}

	r := csv.NewReader(strings.NewReader(string(s3FileData)))
	csvData, err := r.ReadAll()
	if err != nil {
		s.UpdateProcessingStatus(ctx, filename, data.ProcessingStatusFailed, err)
		return err
	}

	err = s.CreateDataPoints(ctx, csvData)
	if err != nil {
		s.UpdateProcessingStatus(ctx, filename, data.ProcessingStatusFailed, err)
		return err
	} else {
		s.logger.Debug("data points created", zap.String("filename", filename))
	}

	s.UpdateProcessingStatus(ctx, filename, data.ProcessingStatusCompleted, nil)
	return nil
}

// UpdateProcessingStatus updates the processing status of a file in the repository.
func (s *DefaultService) UpdateProcessingStatus(ctx context.Context, filename string, status data.ProcessingStatus, err error) error {
	p := data.ProcessingStatusEntity{
		FileName: filename,
		Status:   status,
	}
	if err != nil {
		p.Error = null.NewString(err.Error())
		s.logger.Error("error occurred", zap.String("filename", filename), zap.Error(err))
	}
	return s.repository.UpdateProcessingStatus(ctx, p)
}

// GetProcessingStatus returns the processing status of a file.
func (s *DefaultService) GetProcessingStatus(ctx context.Context, filename string) (*data.ProcessingStatusEntity, error) {
	return s.repository.GetProcessingStatus(ctx, filename)
}

// DeleteStaleProcessingStatus deletes processing status records older than the specified number of days.
func (s *DefaultService) DeleteStaleProcessingStatus(ctx context.Context, days int) error {
	return s.repository.RemoveStaleProcessingStatus(ctx, time.Now().AddDate(0, 0, -days))
}
