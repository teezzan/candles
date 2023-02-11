package ohlc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/teezzan/ohlc/internal/controller/ohlc/data"
	"github.com/teezzan/ohlc/internal/util"
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getFieldTitleIndex(tt.header)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantIsInComplete, got.IsInComplete())
		})
	}
}
