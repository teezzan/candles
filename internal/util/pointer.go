package util

import "time"

// StringPtr return a pointer to the given string.
func StringPtr(s string) *string {
	return &s
}

// StringOr return the derefenced string value if not nil, otherwise fallback.
func StringOr(s *string, fallback string) string {
	if s == nil {
		return fallback
	}
	return *s
}

// BoolPtr return a pointer to the given bool.
func BoolPtr(b bool) *bool {
	return &b
}

// BoolOr return the derefenced bool value if not nil, otherwise fallback.
func BoolOr(b *bool, fallback bool) bool {
	if b == nil {
		return fallback
	}
	return *b
}

// IntPtr return a pointer to the given int.
func IntPtr(i int) *int {
	return &i
}

// IntOr return the derefenced int value if not nil, otherwise fallback.
func IntOr(i *int, fallback int) int {
	if i == nil {
		return fallback
	}
	return *i
}

// Int64Ptr return a pointer to the given int64.
func Int64Ptr(i int64) *int64 {
	return &i
}

// Int64Or return the derefenced int value if not nil, otherwise fallback.
func Int64Or(i *int64, fallback int64) int64 {
	if i == nil {
		return fallback
	}
	return *i
}

// Float64Ptr return a pointer to the given float64.
func Float64Ptr(f float64) *float64 {
	return &f
}

// Float64Or return the derefenced float64 value if not nil, otherwise fallback.
func Float64Or(f *float64, fallback float64) float64 {
	if f == nil {
		return fallback
	}
	return *f
}

// TimePtr return a pointer to the given time.Time.
func TimePtr(t time.Time) *time.Time {
	return &t
}

// TimeOr return the derefenced time.Time value if not nil, otherwise fallback.
func TimeOr(t *time.Time, fallback time.Time) time.Time {
	if t == nil {
		return fallback
	}
	return *t
}
