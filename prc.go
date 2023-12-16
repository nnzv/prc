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
	ErrEmptyRoot   = errors.New("empty root path")
)

// Root is the default path for the "/proc" directory. It is a basic raw string
// that represents the root path without undergoing additional checks, such as
// whether it is mounted. The [Open] function relies on this variable as the root
// path for constructing the proc filepath. If the variable is found to be empty,
// attempting to use it will result in raising an [ErrEmptyRoot] error.
var Root = "/proc"

// File represents an already open file, ready for scanning its content.
type File struct {
	Path string         // Combined [Root] and proc filepath (mostly used for errors)
	f    *os.File       // File handle
	s    *bufio.Scanner // Scanner for reading the file
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

// Open opens a proc file located at the specified path, with the root directory defined by the [Root] variable.
// It checks for errors and validates file properties, returning a [File] struct with the file path
// and a [bufio.Scanner] for reading its content upon success.
func Open(path string) (File, error) {
	if Root == "" {
		return File{}, &ProcError{Op: "open", Err: ErrEmptyRoot}
	}
	path = filepath.Join(Root, path)
	stat, err := os.Stat(path)
	if err != nil {
		return File{}, &ProcError{Err: err} // [fs.PathError] includes the path information.
	}
	if stat.IsDir() {
		return File{}, &ProcError{Op: "open", Path: path, Err: ErrPathIsDir}
	}
	f, err := os.Open(path)
	if err != nil {
		return File{}, err
	}
	buf := new(bytes.Buffer) // file writer
	bts, err := io.ReadAll(io.TeeReader(f, buf))
	if err != nil {
		return File{}, &ProcError{Op: "open", Path: path, Err: err}
	}
	if len(bts) < 1 {
		return File{}, &ProcError{Op: "open", Path: path, Err: ErrFileIsEmpty}
	}
	s := bufio.NewScanner(buf)
	if err := s.Err(); err != nil {
		return File{}, &ProcError{Op: "scan", Path: path, Err: err}
	}
	return File{path, f, s}, nil
}

// ProcError represents an internal operation error
type ProcError struct {
	Op   string // Operator causing the error (optional)
	Path string // File path associated with the error (optional)
	Err  error  // Error details
}

// Error formats the error message.
func (e *ProcError) Error() string {
	if e.Op == "" && e.Path == "" {
		return fmt.Sprintf("proc %s", e.Err)
	}
	if e.Path != "" {
		return fmt.Sprintf("proc %s %s: %s", e.Op, e.Path, e.Err)
	}
	return fmt.Sprintf("proc %s: %s", e.Op, e.Err)
}
