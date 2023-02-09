package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
)

var (
	_ json.Marshaler   = (*Int64)(nil)
	_ json.Unmarshaler = (*Int64)(nil)
	_ sql.Scanner      = (*Int64)(nil)
	_ driver.Valuer    = (*Int64)(nil)
)

// Int64 defines a NULL-able int64 type.
type Int64 struct {
	sql.NullInt64
}

// NewInt64 instantiates a new valid Int64.
func NewInt64(i int64) Int64 {
	return Int64{
		sql.NullInt64{
			Valid: true,
			Int64: i,
		},
	}
}

// NewInt64FromRef sets the value from a pointer if not nil, otherwise
// invalidates.
func NewInt64FromRef(i *int64) Int64 {
	if i == nil {
		return NewInvalidInt64()
	}
	return NewInt64(*i)
}

// NewInvalidInt64 instantiates a new invalid Int64.
func NewInvalidInt64() Int64 {
	return Int64{}
}

// MarshalJSON implements the Marshaler interface.
func (x *Int64) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(x.Int64)
}

// UnmarshalJSON implements the Unmarshaler interface.
func (x *Int64) UnmarshalJSON(data []byte) error {
	var i *int64
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	if i != nil {
		x.Valid = true
		x.Int64 = *i
	} else {
		x.Valid = false
	}
	return nil
}

// Value implements driver.Valuer, will be invoked automatically when written to the db
func (x Int64) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	return int64(x.Int64), nil
}

// Scan implements sql.Scanner, will be invoked automatically when read from the db
func (x *Int64) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*x = Int64{
			NullInt64: sql.NullInt64{
				Valid: false,
			},
		}
	} else {
		*x = Int64{
			NullInt64: sql.NullInt64{
				Valid: true,
				Int64: i.Int64,
			},
		}
	}
	return nil
}

// ValueOr returns the value if valid, otherwise a fallback.
func (x *Int64) ValueOr(fallback int64) int64 {
	if !x.Valid {
		return fallback
	}
	return x.Int64
}

// AsRef returns the value as pointer if valid, otherwise nil.
func (x *Int64) AsRef() *int64 {
	if !x.Valid {
		return nil
	}
	return &x.Int64
}
