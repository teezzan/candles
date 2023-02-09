package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
)

var (
	_ json.Marshaler   = (*Int16)(nil)
	_ json.Unmarshaler = (*Int16)(nil)
	_ sql.Scanner      = (*Int16)(nil)
	_ driver.Valuer    = (*Int16)(nil)
)

// Int16 defines a NULL-able int16 type.
type Int16 struct {
	sql.NullInt16
}

// NewInt16 instantiates a new valid NullInt.
func NewInt16(i int16) Int16 {
	return Int16{
		sql.NullInt16{
			Valid: true,
			Int16: i,
		},
	}
}

// NewInt16FromRef sets the value from a pointer if not nil, otherwise
// invalidates.
func NewInt16FromRef(i *int16) Int16 {
	if i == nil {
		return NewInvalidInt16()
	}
	return NewInt16(*i)
}

// NewInvalidInt16 instantiates a new invalid NullInt.
func NewInvalidInt16() Int16 {
	return Int16{}
}

// MarshalJSON implements the Marshaler interface.
func (x *Int16) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(x.Int16)
}

// UnmarshalJSON implements the Unmarshaler interface.
func (x *Int16) UnmarshalJSON(data []byte) error {
	var i *int16
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	if i != nil {
		x.Valid = true
		x.Int16 = *i
	} else {
		x.Valid = false
	}
	return nil
}

// Value implements driver.Valuer, will be invoked automatically when written to the db
func (x Int16) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	return int16(x.Int16), nil
}

// Scan implements sql.Scanner, will be invoked automatically when read from the db
func (x *Int16) Scan(value interface{}) error {
	var i sql.NullInt16
	if err := i.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*x = Int16{
			NullInt16: sql.NullInt16{
				Valid: false,
			},
		}
	} else {
		*x = Int16{
			NullInt16: sql.NullInt16{
				Valid: true,
				Int16: i.Int16,
			},
		}
	}
	return nil
}

// ValueOr returns the value if valid, otherwise a fallback.
func (x *Int16) ValueOr(fallback int16) int16 {
	if !x.Valid {
		return fallback
	}
	return x.Int16
}

// AsRef returns the value as pointer if valid, otherwise nil.
func (x *Int16) AsRef() *int16 {
	if !x.Valid {
		return nil
	}
	return &x.Int16
}
