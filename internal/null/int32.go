package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
)

var (
	_ json.Marshaler   = (*Int32)(nil)
	_ json.Unmarshaler = (*Int32)(nil)
	_ sql.Scanner      = (*Int32)(nil)
	_ driver.Valuer    = (*Int32)(nil)
)

// Int32 defines a NULL-able int32 type.
type Int32 struct {
	sql.NullInt32
}

// NewInt32 instantiates a new valid Int32.
func NewInt32(i int32) Int32 {
	return Int32{
		sql.NullInt32{
			Valid: true,
			Int32: i,
		},
	}
}

// NewInt32FromRef sets the value from a pointer if not nil, otherwise
// invalidates.
func NewInt32FromRef(i *int32) Int32 {
	if i == nil {
		return NewInvalidInt32()
	}
	return NewInt32(*i)
}

// NewInvalidInt32 instantiates a new invalid Int32.
func NewInvalidInt32() Int32 {
	return Int32{}
}

// MarshalJSON implements the Marshaler interface.
func (x *Int32) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(x.Int32)
}

// UnmarshalJSON implements the Unmarshaler interface.
func (x *Int32) UnmarshalJSON(data []byte) error {
	var b *int32
	if err := json.Unmarshal(data, &b); err != nil {
		return err
	}
	if b != nil {
		x.Valid = true
		x.Int32 = *b
	} else {
		x.Valid = false
	}
	return nil
}

// Value implements driver.Valuer, will be invoked automatically when written to the db
func (x Int32) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	return x.Int32, nil
}

// Scan implements sql.Scanner, will be invoked automatically when read from the db
func (x *Int32) Scan(value interface{}) error {
	var i sql.NullInt32
	if err := i.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*x = Int32{
			NullInt32: sql.NullInt32{
				Valid: false,
			},
		}
	} else {
		*x = Int32{
			NullInt32: sql.NullInt32{
				Valid: true,
				Int32: i.Int32,
			},
		}
	}
	return nil
}

// ValueOr returns the value if valid, otherwise a fallback.
func (x *Int32) ValueOr(fallback int32) int32 {
	if !x.Valid {
		return fallback
	}
	return x.Int32
}

// AsRef returns the value as pointer if valid, otherwise nil.
func (x *Int32) AsRef() *int32 {
	if !x.Valid {
		return nil
	}
	return &x.Int32
}
