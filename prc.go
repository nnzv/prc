// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package prc

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrFileIsEmpty = errors.New("file is empty")
	ErrPathIsDir   = errors.New("path is a directory")
)

// Default path for the "/proc" directory. This path is a simple raw string,
// and no checks, such as whether it is mounted, are performed on it.
var Root = "/proc"

// File represents an already open file, ready for scanning its content.
type File struct {
	Path string         // Combined [Root] and proc filepath (mostly used for errors)
	f    *os.File       // File handle
	s    *bufio.Scanner // Scanner for reading the file
}

// Open opens a proc file located at the specified path, with the root directory defined by the [Root] variable.
// It checks for errors and validates file properties, returning a [File] struct with the file path
// and a [bufio.Scanner] for reading its content upon success.
func Open(path string) (*File, error) {
	p := filepath.Join(Root, path)
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	stat, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("proc %q", err) // [fs.PathError] includes the path information.
	}
	if stat.IsDir() {
		return nil, fmt.Errorf("proc %s: %q", p, ErrPathIsDir)
	}
	buf := new(bytes.Buffer) // file writer
	tee := io.TeeReader(f, buf)
	bts, err := io.ReadAll(tee)
	if err != nil {
		return nil, err
	}
	if len(bts) < 1 {
		return nil, fmt.Errorf("proc %s: %q", p, ErrFileIsEmpty)
	}
	sc := bufio.NewScanner(buf)
	if err := sc.Err(); err != nil {
		f.Close()
		return nil, fmt.Errorf("proc %s: %q", p, err)
	}
	return &File{p, f, sc}, nil
}

// Close closes the /proc file by closing its [os.File] handle.
func (f *File) Close() error { return f.f.Close() }

// Scan scans the /proc file by uing its [bufio.Scanner] handle.
func (f *File) Scan() bool { return f.s.Scan() }

// SplitWords configures the scanner to split words using [bufio.ScanWords].
func (f *File) SplitWords() { f.s.Split(bufio.ScanWords) }

// ScanFields scans the current line of the file, splitting it into fields using whitespace.
// Returns a slice of strings representing the fields.
func (f *File) ScanFields() []string { return strings.Fields(f.s.Text()) }

// ParseError represents a parsing error, including the field name, file path, and the encountered error.
type ParseError struct {
	Path  string // File path causing the parsing error
	Field string // Field name causing the parsing error
	Err   error  // Parsing error details
}

// Error formats the error message.
func (e *ParseError) Error() string {
	return fmt.Sprintf("parsing %s in %s: %q", e.Field, e.Path, e.Err)
}
