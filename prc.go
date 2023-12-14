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
func Open(path string) (File, error) {
	path = filepath.Join(Root, path)
	f, err := os.Open(path)
	if err != nil {
		return File{}, err
	}
	stat, err := f.Stat()
	if err != nil {
		return File{}, &ProcError{Op: "open", Err: err} // [fs.PathError] includes the path information.
	}
	if stat.IsDir() {
		return File{}, &ProcError{Op: "open", Path: path, Err: ErrPathIsDir}
	}
	buf := new(bytes.Buffer) // file writer
	bts, err := io.ReadAll(io.TeeReader(f, buf))
	if err != nil {
		return File{}, err
	}
	if len(bts) < 1 {
		return File{}, &ProcError{Op: "open", Path: path, Err: ErrFileIsEmpty}
	}
	s := bufio.NewScanner(buf)
	if err := s.Err(); err != nil {
		f.Close()
		return File{}, &ProcError{Op: "scan", Path: path, Err: err}
	}
	return File{path, f, s}, nil
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

// ProcError represents an internal operation error
type ProcError struct {
	Op   string // Operator causing the error
	Path string // File path associated with the error (optional)
	Err  error  // Error details
}

// Error formats the error message.
func (e *ProcError) Error() string {
	var b strings.Builder
	b.WriteString("proc %s")
	if e.Path != "" {
		b.WriteString(" %s")
	}
	b.WriteString(": %s")
	return fmt.Sprintf(b.String(), e.Op, e.Path, e.Err)
}
