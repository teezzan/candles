package util

import (
	"os"
	"strconv"
	"strings"
)

// Keys is a list of keys that were retrieved from environment.
var Keys = []string{}

// GetString retrieves an environment variable and parses it as string.
func GetString(key, fallback string) string {
	Keys = append(Keys, key)
	if v, ok := os.LookupEnv(key); ok {
		return strings.TrimSpace(v)
	}
	return fallback
}

// GetInt retrieves an environment variable and parses it as int.
func GetInt(key string, fallback int) int {
	Keys = append(Keys, key)
	if v, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(v)
		if err == nil {
			return i
		}
	}
	return fallback
}

// GetBool retrieves an environment variable and parses it as bool.
//
// Note, only `true` is a valid true value.
func GetBool(key string, fallback bool) bool {
	Keys = append(Keys, key)
	if v, ok := os.LookupEnv(key); ok {
		return v == "true"
	}
	return fallback
}
