package data

import (
	"time"

	"github.com/teezzan/candles/internal/null"
)

// Project defines the ohlc project DTO
type OHLC struct {
	Time   int64   `json:"unix"`
	Symbol string  `json:"symbol"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Low    float64 `json:"low"`
}

// OHLCEntity defines the ohlc data.
type OHLCEntity struct {
	ID     int64     `db:"id"`
	Time   time.Time `db:"time"`
	Symbol string    `db:"symbol"`
	Open   float64   `db:"open"`
	High   float64   `db:"high"`
	Close  float64   `db:"close"`
	Low    float64   `db:"low"`
}

// OHLCEntity converts OHLCEntity to OHLC
func (p *OHLCEntity) ToOHLC() OHLC {
	return OHLC{
		Time:   p.Time.Unix(),
		Symbol: p.Symbol,
		Open:   p.Open,
		High:   p.High,
		Close:  p.Close,
		Low:    p.Low,
	}
}

// IsInComplete returns true if the OHLCEntity is incomplete.
func (p *OHLCEntity) IsInComplete() bool {
	return p.Time.IsZero() || p.Symbol == "" || p.Open == 0 || p.High == 0 || p.Close == 0 || p.Low == 0
}

// OHLCFieldName defines the OHLC Field name.
type OHLCFieldName string

// Stringer implements the Stringer interface.
func (c OHLCFieldName) String() string {
	return string(c)
}

// FieldIndex defines the OHLC Field index.
type FieldIndex struct {
	Name  OHLCFieldName
	Index *int
}

// IsEmptyIndex returns true if the index is empty.
func (c *FieldIndex) IsEmptyIndex() bool {
	return c.Index == nil
}

// OHLCFieldIndexes defines the OHLC Field indexes.
type OHLCFieldIndexes struct {
	Open   FieldIndex
	High   FieldIndex
	Low    FieldIndex
	Close  FieldIndex
	Symbol FieldIndex
	Unix   FieldIndex
}

// IsInComplete returns true if the OHLCFieldIndexes is incomplete.
func (c *OHLCFieldIndexes) IsInComplete() bool {
	return c.Open.IsEmptyIndex() || c.High.IsEmptyIndex() || c.Low.IsEmptyIndex() || c.Close.IsEmptyIndex() || c.Symbol.IsEmptyIndex() || c.Unix.IsEmptyIndex()
}

// GetOHLCRequest defines the get ohlc request.
type GetOHLCRequest struct {
	Symbol     string     `form:"symbol"`
	StartTime  int64      `form:"from"`
	EndTime    null.Int64 `form:"to"`
	PageNumber null.Int   `form:"page"`
	PageSize   null.Int   `form:"page_size"`
}

// GetOHLCResponse defines the get ohlc response.
type GetOHLCResponse struct {
	DataPoints []OHLC `json:"data"`
	Page       int    `json:"page"`
}

// GeneratePresignedURLResponse defines the generate presigned url response.
type GeneratePresignedURLResponse struct {
	URL      string `json:"url"`
	Filename string `json:"filename"`
}
