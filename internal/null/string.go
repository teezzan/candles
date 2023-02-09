package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
)

var (
	_ json.Marshaler   = (*String)(nil)
	_ json.Unmarshaler = (*String)(nil)
	_ sql.Scanner      = (*String)(nil)
	_ driver.Valuer    = (*String)(nil)
)

// String defines a NULL-able string type.
type String struct {
	sql.NullString
}

// NewString instantiates a new valid String.
func NewString(s string) String {
	return String{
		sql.NullString{
			Valid:  true,
			String: s,
		},
	}
}

// NewStringFromRef sets the value from a pointer if not nil, otherwise
// invalidates.
func NewStringFromRef(s *string) String {
	if s == nil {
		return NewInvalidString()
	}
	return NewString(*s)
}

// NewInvalidString instantiates a new invalid String.
func NewInvalidString() String {
	return String{}
}

// MarshalJSON implements the Marshaler interface.
func (x *String) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(x.String)
}

// UnmarshalJSON implements the Unmarshaler interface.
func (x *String) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		x.Valid = true
		x.String = *s
	} else {
		x.Valid = false
	}
	return nil
}

// Value implements driver.Valuer, will be invoked automatically when written to the db
func (x String) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	return x.String, nil
}

// Scan implements sql.Scanner, will be invoked automatically when read from the db
func (x *String) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*x = String{
			sql.NullString{
				Valid: false,
			},
		}
	} else {
		*x = String{
			sql.NullString{
				Valid:  true,
				String: s.String,
			},
		}
	}
	return nil
}

// ValueOr returns the value if valid, otherwise a fallback.
func (x *String) ValueOr(fallback string) string {
	if !x.Valid {
		return fallback
	}
	return x.String
}

// AsRef returns the value as pointer if valid, otherwise nil.
func (x *String) AsRef() *string {
	if !x.Valid {
		return nil
	}
	return &x.String
}

// IsEmptyString returns true if the string is empty.
func (x *String) IsEmptyString() bool {
	return x.Valid && x.String == ""
}
