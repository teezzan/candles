package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
)

var (
	_ json.Marshaler   = (*Float64)(nil)
	_ json.Unmarshaler = (*Float64)(nil)
	_ sql.Scanner      = (*Float64)(nil)
	_ driver.Valuer    = (*Float64)(nil)
)

// Float64 defines a NULL-able float64 type.
type Float64 struct {
	sql.NullFloat64
}

// NewFloat64 instantiates a new valid NullFloat.
func NewFloat64(f float64) Float64 {
	return Float64{
		sql.NullFloat64{
			Valid:   true,
			Float64: f,
		},
	}
}

// NewFloat64FromRef sets the value from a pointer if not nil, otherwise
// invalidates.
func NewFloat64FromRef(f *float64) Float64 {
	if f == nil {
		return NewInvalidFloat64()
	}
	return NewFloat64(*f)
}

// NewInvalidFloat64 instantiates a new invalid NullFloat.
func NewInvalidFloat64() Float64 {
	return Float64{}
}

// MarshalJSON implements the Marshaler interface.
func (x *Float64) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(x.Float64)
}

// UnmarshalJSON implements the Unmarshaler interface.
func (x *Float64) UnmarshalJSON(data []byte) error {
	var f *float64
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}
	if f != nil {
		x.Valid = true
		x.Float64 = *f
	} else {
		x.Valid = false
	}
	return nil
}

// Value implements driver.Valuer, will be invoked automatically when written to the db
func (x Float64) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	return x.Float64, nil
}

// Scan implements sql.Scanner, will be invoked automatically when read from the db
func (x *Float64) Scan(value interface{}) error {
	var f sql.NullFloat64
	if err := f.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*x = Float64{
			sql.NullFloat64{
				Valid: false,
			},
		}
	} else {
		*x = Float64{
			sql.NullFloat64{
				Valid:   true,
				Float64: f.Float64,
			},
		}
	}
	return nil
}

// ValueOr returns the value if valid, otherwise a fallback.
func (x *Float64) ValueOr(fallback float64) float64 {
	if !x.Valid {
		return fallback
	}
	return x.Float64
}

// AsRef returns the value as pointer if valid, otherwise nil.
func (x *Float64) AsRef() *float64 {
	if !x.Valid {
		return nil
	}
	return &x.Float64
}
