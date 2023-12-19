// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package internal

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ParseInt parses an integer from a string with the specified base and bit size.
// It returns the parsed integer and any error encountered during parsing.
func ParseInt(from string, base int, bitSize int) (int64, error) {
	to, err := strconv.ParseInt(from, base, bitSize)
	if err != nil {
		return 0, err
	}
	return to, nil
}

// ParseFloat parses a float64 from a string.
// It returns the parsed float64 and any error encountered during parsing.
func ParseFloat(from string) (float64, error) {
	to, err := strconv.ParseFloat(from, 64)
	if err != nil {
		return 0, err
	}
	return to, nil
}

// ParseDuration parses a time.Duration from a string.
// It internally uses ParseFloat to convert the string to a float64,
// and then converts it to a time.Duration in seconds.
// It returns the parsed time.Duration and any error encountered during parsing.
func ParseDuration(from string) (time.Duration, error) {
	to, err := ParseFloat(from)
	if err != nil {
		return 0, err
	}
	return time.Duration(to * float64(time.Second)), nil
}

// ParseUint64 parses a uint64 from a string with the specified base and bit size.
// It returns the parsed uint64 and any error encountered during parsing.
func ParseUint64(from string, base int, bitSize int) (uint64, error) {
	to, err := strconv.ParseUint(from, base, bitSize)
	if err != nil {
		return 0, err
	}
	return to, nil
}

// ParseCommaList parses a comma-separated string and returns a slice of strings.
func ParseCommaList(from string) []string {
	if from[len(from)-1:] == "," {
		return []string{from[:len(from)-1]}
	}
	return strings.Split(from, ",")
}

// ParseError represents a parsing error
type ParseError struct {
	Path  string // File path causing the parsing error
	Field string // Field name causing the parsing error
	Err   error  // Parsing error details
}

// Error formats the error message.
func (e *ParseError) Error() string {
	return fmt.Sprintf("proc parse %s (%s): %s", e.Path, e.Field, e.Err)
}
