package ohlc

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/teezzan/ohlc/internal/client/s3"
	"github.com/teezzan/ohlc/internal/client/sqs"
	"github.com/teezzan/ohlc/internal/config"
	"github.com/teezzan/ohlc/internal/controller/ohlc/data"
	"github.com/teezzan/ohlc/internal/controller/ohlc/repository"
	"github.com/teezzan/ohlc/internal/util"
	"go.uber.org/zap"
)

func Test_extractDataPoint(t *testing.T) {
	tests := []struct {
		name         string
		row          []string
		fieldIndexes data.OHLCFieldIndexes
		want         *data.OHLCEntity
		wantErr      bool
	}{
		{
			name: "valid row",
			row: []string{
				"1610000000",
				"BTC/USD",
				"100",
				"200",
				"50",
				"150",
			},
			fieldIndexes: getFieldTitleIndex([]string{"UNIX", "SYMBOL", "OPEN", "HIGH", "LOW", "CLOSE"}),
			want: &data.OHLCEntity{
				Time:   time.Unix(1610000000, 0),
				Symbol: "BTC/USD",
				Open:   100,
				High:   200,
				Low:    50,
				Close:  150,
			},
			wantErr: false,
		},
		{
			name: "invalid row",
			row: []string{
				"1610000000",
				"BTC/USD",
				"a100",
				"b200",
				"c50",
				"d150",
			},
			fieldIndexes: getFieldTitleIndex([]string{"UNIX", "SYMBOL", "OPEN", "HIGH", "LOW", "CLOSE"}),
			want:         nil,
			wantErr:      true,
		},
		{
			name: "incomplete row",
			row: []string{
				"1610000000",
				"BTC/USD",
				"100",
				"200",
			},
			fieldIndexes: getFieldTitleIndex([]string{"UNIX", "SYMBOL", "OPEN", "HIGH", "LOW", "CLOSE"}),
			want:         nil,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractDataPoint(tt.row, tt.fieldIndexes)
			require.Equal(t, tt.wantErr, err != nil)
			if !tt.wantErr {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_getFieldTitleIndex(t *testing.T) {
	tests := []struct {
		name             string
		header           []string
		want             data.OHLCFieldIndexes
		wantIsInComplete bool
	}{
		{
			name:   "valid header",
			header: []string{"UNIX", "SYMBOL", "OPEN", "HIGH", "LOW", "CLOSE"},
			want: data.OHLCFieldIndexes{
				Unix: data.FieldIndex{
					Index: util.IntPtr(0),
					Name:  "UNIX",
				},
				Symbol: data.FieldIndex{
					Index: util.IntPtr(1),
					Name:  "SYMBOL",
				},
				Open: data.FieldIndex{
					Index: util.IntPtr(2),
					Name:  "OPEN",
				},
				High: data.FieldIndex{
					Index: util.IntPtr(3),
					Name:  "HIGH",
				},
				Low: data.FieldIndex{
					Index: util.IntPtr(4),
					Name:  "LOW",
				},
				Close: data.FieldIndex{
					Index: util.IntPtr(5),
					Name:  "CLOSE",
				},
			},
			wantIsInComplete: false,
		},
		{
			name:   "valid header in different order",
			header: []string{"HIGH", "LOW", "UNIX", "SYMBOL", "OPEN", "CLOSE"},
			want: data.OHLCFieldIndexes{
				High: data.FieldIndex{
					Index: util.IntPtr(0),
					Name:  "HIGH",
				},
				Low: data.FieldIndex{
					Index: util.IntPtr(1),
					Name:  "LOW",
				},
				Unix: data.FieldIndex{
					Index: util.IntPtr(2),
					Name:  "UNIX",
				},
				Symbol: data.FieldIndex{
					Index: util.IntPtr(3),
					Name:  "SYMBOL",
				},
				Open: data.FieldIndex{
					Index: util.IntPtr(4),
					Name:  "OPEN",
				},
				Close: data.FieldIndex{
					Index: util.IntPtr(5),
					Name:  "CLOSE",
				},
			},
			wantIsInComplete: false,
		},
		{
			name:   "valid header with extra fields",
			header: []string{"HIGH", "LOW", "UNIX", "SYMBOL", "OPEN", "CLOSE", "EXTRA"},
			want: data.OHLCFieldIndexes{
				High: data.FieldIndex{
					Index: util.IntPtr(0),
					Name:  "HIGH",
				},
				Low: data.FieldIndex{
					Index: util.IntPtr(1),
					Name:  "LOW",
				},
				Unix: data.FieldIndex{
					Index: util.IntPtr(2),
					Name:  "UNIX",
				},
				Symbol: data.FieldIndex{
					Index: util.IntPtr(3),
					Name:  "SYMBOL",
				},
				Open: data.FieldIndex{
					Index: util.IntPtr(4),
					Name:  "OPEN",
				},
				Close: data.FieldIndex{
					Index: util.IntPtr(5),
					Name:  "CLOSE",
				},
			},
			wantIsInComplete: false,
		},
		{
			name:             "invalid header",
			header:           []string{"EXTRA", "VALID", "TITLE", "RANDOM"},
			want:             data.DefaultOHLCFieldIndexes,
			wantIsInComplete: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getFieldTitleIndex(tt.header)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantIsInComplete, got.IsInComplete())
		})
	}
}

func TestDefaultService_CreateDataPoints(t *testing.T) {
	tests := []struct {
		name                     string
		discardInCompleteRow     bool
		repository               repository.RepositoryMock
		dataPoints               [][]string
		wantErr                  bool
		InsertDataPointsCallsNum int
	}{
		{
			name: "valid data points",
			repository: repository.RepositoryMock{
				InsertDataPointsFunc: func(ctx context.Context, rows []data.OHLCEntity) error {
					return nil
				},
			},
			dataPoints: [][]string{
				{"UNIX", "SYMBOL", "OPEN", "HIGH", "LOW", "CLOSE"},
				{"1610000000", "BTC", "100", "200", "50", "150"},
				{"1610000001", "BTC", "150", "250", "100", "200"},
			},
			wantErr:                  false,
			InsertDataPointsCallsNum: 1,
		},
		{
			name:       "invalid csv row header",
			repository: repository.RepositoryMock{},
			dataPoints: [][]string{
				{"EXTRA", "VALID", "TITLE", "RANDOM", "CLOSE"},
				{"1610000000", "BTC", "100", "200", "50", "150"},
				{"1610000001", "BTC", "150", "250", "100", "200"},
			},
			wantErr:                  true,
			InsertDataPointsCallsNum: 0,
		},
		{
			name:       "invalid csv row with discardInCompleteRow to be false",
			repository: repository.RepositoryMock{},
			dataPoints: [][]string{
				{"UNIX", "SYMBOL", "OPEN", "HIGH", "LOW", "CLOSE"},
				{"1610000000", "BTC", "a100", "a200", "a50", "a150"},
				{"1610000001", "BTC", "150", "250", "100", "200"},
			},
			wantErr:                  true,
			InsertDataPointsCallsNum: 0,
		},
		{
			name:                 "invalid csv row with discardInCompleteRow to be true",
			discardInCompleteRow: true,
			repository: repository.RepositoryMock{
				InsertDataPointsFunc: func(ctx context.Context, rows []data.OHLCEntity) error {
					return nil
				},
			},
			dataPoints: [][]string{
				{"UNIX", "SYMBOL", "OPEN", "HIGH", "LOW", "CLOSE"},
				{"1610000000", "BTC", "a100", "a200", "a50", "a150"},
				{"1610000001", "BTC", "150", "250", "100", "200"},
			},
			wantErr:                  false,
			InsertDataPointsCallsNum: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				ctx           = context.Background()
				logger        = zap.NewNop()
				mockS3Client  = &s3.ClientMock{}
				mockSQSClient = &sqs.ClientMock{}
			)
			conf := config.Init()
			conf.OHLCConfig.DiscardInCompleteRow = tt.discardInCompleteRow

			s := NewService(logger, &tt.repository, mockS3Client, mockSQSClient, conf.OHLCConfig)
			err := s.CreateDataPoints(ctx, tt.dataPoints)
			require.Equal(t, tt.wantErr, err != nil)
			assert.Len(t, tt.repository.InsertDataPointsCalls(), tt.InsertDataPointsCallsNum)
		})
	}
}

func TestDefaultService_GeneratePreSignedURL(t *testing.T) {
	tests := []struct {
		name     string
		s3Client s3.ClientMock
		wantErr  bool
	}{
		{
			name: "valid s3 client response",
			s3Client: s3.ClientMock{
				GeneratePresignedURLFunc: func(ctx context.Context, key string) (string, error) {
					return "https://test.com", nil
				},
			},
			wantErr: false,
		},
		{
			name: "s3 client respond with error",
			s3Client: s3.ClientMock{
				GeneratePresignedURLFunc: func(ctx context.Context, key string) (string, error) {
					return "", errors.New("test error")
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				ctx            = context.Background()
				logger         = zap.NewNop()
				mockSQSClient  = &sqs.ClientMock{}
				mockRepository = &repository.RepositoryMock{}
			)
			conf := config.Init()

			s := NewService(logger, mockRepository, &tt.s3Client, mockSQSClient, conf.OHLCConfig)
			got, err := s.GeneratePreSignedURL(ctx)
			require.Equal(t, tt.wantErr, err != nil)
			if !tt.wantErr {
				assert.NotEmpty(t, got)
				assert.Len(t, tt.s3Client.GeneratePresignedURLCalls(), 1)
			}
		})
	}
}

var validCSV = `UNIX,SYMBOL,OPEN,HIGH,LOW,CLOSE
1610000000,BTC,100,200,50,150
1610000001,BTC,150,250,100,200
`
var invalidCSV = `UNIX,SYMBOL,TOP,HIGH,LOW,CLOSE
1610000000,BTC,100,200,50,150
1610000001,BTC,150,250,100,200
`

func TestDefaultService_DownloadAndProcessCSV(t *testing.T) {
	tests := []struct {
		name     string
		s3Client s3.ClientMock
		wantErr  bool
	}{
		{
			name: "valid CSV file without error",
			s3Client: s3.ClientMock{
				DownloadLargeObjectFunc: func(ctx context.Context, objectKey string) ([]byte, error) {
					return []byte(validCSV), nil
				},
			},
			wantErr: false,
		},
		{
			name: "invalid CSV file without error",
			s3Client: s3.ClientMock{
				DownloadLargeObjectFunc: func(ctx context.Context, objectKey string) ([]byte, error) {
					return []byte(invalidCSV), nil
				},
			},
			wantErr: true,
		},
		{
			name: "no CSV file with error",
			s3Client: s3.ClientMock{
				DownloadLargeObjectFunc: func(ctx context.Context, objectKey string) ([]byte, error) {
					return nil, errors.New("test error")
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				ctx            = context.Background()
				logger         = zap.NewNop()
				mockSQSClient  = &sqs.ClientMock{}
				mockRepository = &repository.RepositoryMock{
					InsertDataPointsFunc: func(ctx context.Context, rows []data.OHLCEntity) error {
						return nil
					},
				}
			)
			conf := config.Init()

			s := NewService(logger, mockRepository, &tt.s3Client, mockSQSClient, conf.OHLCConfig)
			err := s.DownloadAndProcessCSV(ctx, "test")
			require.Equal(t, tt.wantErr, err != nil)
			if !tt.wantErr {
				assert.Len(t, tt.s3Client.DownloadLargeObjectCalls(), 1)
				assert.Len(t, mockRepository.InsertDataPointsCalls(), 1)

			}
		})
	}
}

func TestDefaultService_GetAndProcessSQSMessage(t *testing.T) {
	tests := []struct {
		name      string
		sqsClient sqs.ClientMock
		wantErr   bool
	}{
		{
			name: "valid filenames without error",
			sqsClient: sqs.ClientMock{
				GetFilenamesFromMessagesFunc: func(ctx context.Context) ([]string, error) {
					return []string{"test.csv"}, nil
				},
			},
			wantErr: false,
		},
		{
			name: "No file without error",
			sqsClient: sqs.ClientMock{
				GetFilenamesFromMessagesFunc: func(ctx context.Context) ([]string, error) {
					return nil, nil
				},
			},
			wantErr: false,
		},
		{
			name: "no file with error",
			sqsClient: sqs.ClientMock{
				GetFilenamesFromMessagesFunc: func(ctx context.Context) ([]string, error) {
					return nil, errors.New("test error")
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				ctx          = context.Background()
				logger       = zap.NewNop()
				mockS3Client = &s3.ClientMock{
					DownloadLargeObjectFunc: func(ctx context.Context, objectKey string) ([]byte, error) {
						return []byte(validCSV), nil
					},
				}
				mockRepository = &repository.RepositoryMock{
					InsertDataPointsFunc: func(ctx context.Context, rows []data.OHLCEntity) error {
						return nil
					},
				}
			)
			conf := config.Init()

			s := NewService(logger, mockRepository, mockS3Client, &tt.sqsClient, conf.OHLCConfig)
			err := s.GetAndProcessSQSMessage(ctx)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}
