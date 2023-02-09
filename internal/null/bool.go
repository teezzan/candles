package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
)

var (
	_ json.Marshaler   = (*Bool)(nil)
	_ json.Unmarshaler = (*Bool)(nil)
	_ sql.Scanner      = (*Bool)(nil)
	_ driver.Valuer    = (*Bool)(nil)
)

// Bool defines a NULL-able bool type.
type Bool struct {
	sql.NullBool
}

// NewBool instantiates a new valid Bool.
func NewBool(b bool) Bool {
	return Bool{
		sql.NullBool{
			Bool:  b,
			Valid: true,
		},
	}
}

// NewBoolFromRef sets the value from a pointer if not nil, otherwise
// invalidates.
func NewBoolFromRef(b *bool) Bool {
	if b == nil {
		return NewInvalidBool()
	}
	return NewBool(*b)
}

// NewInvalidBool instantiates a new invalid Bool.
func NewInvalidBool() Bool {
	return Bool{}
}

// MarshalJSON implements the Marshaler interface.
func (x *Bool) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		// Alwayes return false if nil
		return json.Marshal(false)
	}
	return json.Marshal(x.Bool)
}

// UnmarshalJSON implements the Unmarshaler interface.
func (x *Bool) UnmarshalJSON(data []byte) error {
	var b *bool
	if err := json.Unmarshal(data, &b); err != nil {
		return err
	}
	if b != nil {
		x.Valid = true
		x.Bool = *b
	} else {
		x.Valid = false
	}
	return nil
}

// Value implements driver.Valuer, will be invoked automatically when written to the db
func (x Bool) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	return x.Bool, nil
}

// Scan implements sql.Scanner, will be invoked automatically when read from the db
func (x *Bool) Scan(value interface{}) error {
	var b sql.NullBool
	if err := b.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*x = Bool{
			sql.NullBool{
				Valid: false,
			},
		}
	} else {
		*x = Bool{
			sql.NullBool{
				Valid: true,
				Bool:  b.Bool,
			},
		}
	}
	return nil
}

// ValueOr returns the value if valid, otherwise a fallback.
func (x *Bool) ValueOr(fallback bool) bool {
	if !x.Valid {
		return fallback
	}
	return x.Bool
}

// AsRef returns the value as pointer if valid, otherwise nil.
func (x *Bool) AsRef() *bool {
	if !x.Valid {
		return nil
	}
	return &x.Bool
}
