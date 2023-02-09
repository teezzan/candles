package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"time"
)

var (
	_ json.Marshaler   = (*Time)(nil)
	_ json.Unmarshaler = (*Time)(nil)
	_ sql.Scanner      = (*Time)(nil)
	_ driver.Valuer    = (*Time)(nil)
)

// Time defines a NULL-able time.Time type.
type Time struct {
	sql.NullTime
}

// NewTime instantiates a new valid Time.
func NewTime(t time.Time) Time {
	return Time{
		sql.NullTime{
			Valid: true,
			Time:  t,
		},
	}
}

// NewTimeFromRef sets the value from a pointer if not nil, otherwise
// invalidates.
func NewTimeFromRef(t *time.Time) Time {
	if t == nil {
		return NewInvalidTime()
	}
	return NewTime(*t)
}

// NewInvalidTime instantiates a new invalid Time.
func NewInvalidTime() Time {
	return Time{}
}

// MarshalJSON implements the Marshaler interface.
func (x *Time) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(x.Time)
}

// ToUnix is a convenience function converting a potential time.Time to unix
// timestamp.
func (x *Time) ToUnix() *int64 {
	if !x.Valid {
		return nil
	}
	unix := x.Time.Unix()
	return &unix
}

// UnmarshalJSON implements the Unmarshaler interface.
func (x *Time) UnmarshalJSON(data []byte) error {
	var t *time.Time
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	if t != nil {
		x.Valid = true
		x.Time = *t
	} else {
		x.Valid = false
	}
	return nil
}

// Value implements driver.Valuer, will be invoked automatically when written to the db
func (x Time) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	return x.Time, nil
}

// Scan implements sql.Scanner, will be invoked automatically when read from the db
func (x *Time) Scan(value interface{}) error {
	var t sql.NullTime
	if err := t.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*x = Time{
			NullTime: sql.NullTime{
				Valid: false,
				Time:  t.Time,
			},
		}
	} else {
		*x = Time{
			NullTime: sql.NullTime{
				Valid: true,
				Time:  t.Time,
			},
		}
	}
	return nil
}

// ValueOr returns the value if valid, otherwise a fallback.
func (x *Time) ValueOr(fallback time.Time) time.Time {
	if !x.Valid {
		return fallback
	}
	return x.Time
}

// AsRef returns the value as pointer if valid, otherwise nil.
func (x *Time) AsRef() *time.Time {
	if !x.Valid {
		return nil
	}
	return &x.Time
}
