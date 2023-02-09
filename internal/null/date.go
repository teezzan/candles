package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"time"
)

var (
	_ json.Marshaler   = (*Date)(nil)
	_ json.Unmarshaler = (*Date)(nil)
	_ sql.Scanner      = (*Date)(nil)
	_ driver.Valuer    = (*Date)(nil)
)

const layout = "2006-01-02"

// Date defines a NULL-able date type.
type Date struct {
	sql.NullTime
}

// NewDate instantiates a new valid Date.
func NewDate(t time.Time) Date {
	return Date{
		sql.NullTime{
			Valid: true,
			Time:  t,
		},
	}
}

// NewDateFromRef sets the value from a pointer if not nil, otherwise
// invalidates.
func NewDateFromRef(t *time.Time) Date {
	if t == nil {
		return NewInvalidDate()
	}
	return NewDate(*t)
}

// NewInvalidDate instantiates a new invalid Date.
func NewInvalidDate() Date {
	return Date{}
}

// MarshalJSON implements the Marshaler interface.
func (x *Date) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(x.Time.Format(layout))
}

// UnmarshalJSON implements the Unmarshaler interface.
func (x *Date) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		x.Valid = true
		t, _ := time.Parse(layout, *s)
		x.Time = t
	} else {
		x.Valid = false
	}
	return nil
}

// Value implements driver.Valuer, will be invoked automatically when written to the db
func (x Date) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	return x.Time, nil
}

// Scan implements sql.Scanner, will be invoked automatically when read from the db
func (x *Date) Scan(value interface{}) error {
	var t sql.NullTime
	if err := t.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*x = Date{
			sql.NullTime{
				Valid: false,
			},
		}
	} else {
		*x = Date{
			sql.NullTime{
				Valid: true,
				Time:  t.Time,
			},
		}
	}
	return nil
}

// ValueOr returns the value if valid, otherwise a fallback.
func (x *Date) ValueOr(fallback time.Time) time.Time {
	if !x.Valid {
		return fallback
	}
	return x.Time
}

// AsRef returns the value as pointer if valid, otherwise nil.
func (x *Date) AsRef() *time.Time {
	if !x.Valid {
		return nil
	}
	return &x.Time
}
