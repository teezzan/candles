package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
)

var (
	_ json.Marshaler   = (*Int)(nil)
	_ json.Unmarshaler = (*Int)(nil)
	_ sql.Scanner      = (*Int)(nil)
	_ driver.Valuer    = (*Int)(nil)
)

// Int defines a NULL-able int type.
type Int struct {
	sql.NullInt64
}

// NewInt instantiates a new valid Int.
func NewInt(i int) Int {
	return Int{
		sql.NullInt64{
			Valid: true,
			Int64: int64(i),
		},
	}
}

// NewIntFromRef sets the value from a pointer if not nil, otherwise
// invalidates.
func NewIntFromRef(i *int) Int {
	if i == nil {
		return NewInvalidInt()
	}
	return NewInt(*i)
}

// NewInvalidInt instantiates a new invalid Int.
func NewInvalidInt() Int {
	return Int{}
}

// MarshalJSON implements the Marshaler interface.
func (x *Int) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(x.Int64)
}

// UnmarshalJSON implements the Unmarshaler interface.
func (x *Int) UnmarshalJSON(data []byte) error {
	var b *int
	if err := json.Unmarshal(data, &b); err != nil {
		return err
	}
	if b != nil {
		x.Valid = true
		x.Int64 = int64(*b)
	} else {
		x.Valid = false
	}
	return nil
}

// Value implements driver.Valuer, will be invoked automatically when written to
// the db.
func (x Int) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	return x.Int64, nil
}

// Scan implements sql.Scanner, will be invoked automatically when read from the
// db.
func (x *Int) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*x = Int{
			NullInt64: sql.NullInt64{
				Valid: false,
			},
		}
	} else {
		*x = Int{
			NullInt64: sql.NullInt64{
				Valid: true,
				Int64: i.Int64,
			},
		}
	}
	return nil
}

// ValueOr returns the value if valid, otherwise a fallback.
func (x *Int) ValueOr(fallback int) int {
	if !x.Valid {
		return fallback
	}
	return int(x.Int64)
}

// AsRef returns the value as pointer if valid, otherwise nil.
func (x *Int) AsRef() *int {
	if !x.Valid {
		return nil
	}
	i := int(x.Int64)
	return &i
}
