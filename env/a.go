package env

import (
	"os"
	"strconv"
)

// GetenvAsString ...
func GetenvAsString(name string, dflt string) string {
	v := os.Getenv(name)
	if v == "" {
		return dflt
	}
	return v
}

// GetenvAsInt ...
func GetenvAsInt(name string, dflt int) int {
	v := os.Getenv(name)
	ret, err := strconv.Atoi(v)
	if err != nil {
		return dflt
	}
	return ret
}

// GetenvAsBool ...
func GetenvAsBool(name string, dflt bool) bool {
	v := os.Getenv(name)
	ret, err := strconv.ParseBool(v)
	if err != nil {
		return dflt
	}
	return ret
}