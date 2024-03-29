// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package internal

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gitlab.com/nzv/prc"
)

var (
	ErrFileIsEmpty = errors.New("file is empty")
	ErrPathIsDir   = errors.New("path is a directory")
	ErrEmptyRoot   = errors.New("empty root path")
)

// File represents an already open file, ready for scanning its content.
type File struct {
	Path string         // Combined [prc.Root] and proc filepath (mostly used for errors)
	f    *os.File       // File handle
	s    *bufio.Scanner // Scanner for reading the file
}

// Close closes the /proc file by closing its [os.File] handle.
func (f *File) Close() error { return f.f.Close() }

// Scan scans the /proc file by uing its [bufio.Scanner] handle.
func (f *File) Scan() bool { return f.s.Scan() }

// SplitWords configures the scanner to split words using [bufio.ScanWords].
func (f *File) SplitWords() { f.s.Split(bufio.ScanWords) }

func (f *File) Text() string { return f.s.Text() }

// ScanFields scans the current line of the file, splitting it into fields using whitespace.
// Returns a slice of strings representing the fields.
func (f *File) ScanFields() []string { return strings.Fields(f.Text()) }

// Open opens a proc file specified by the given path and returns a [File] and an error, if any.
// In the presence of any error, the returned file will be an empty instance of [File]. It
// returns [ErrEmptyRoot] if the global variable [Root] is empty. The full path is constructed by
// joining [Root] and the provided path. The function checks if the path exists and is not a directory;
// otherwise, it returns [ErrPathIsDir]. Since proc files lack a true size, the function reads the
// file bytes directly instead of using the "Size" method of [io/fs.FileInfo]. If the content is less
// than 1 byte, returns [ErrFileIsEmpty].
func Open(path string) (File, error) {
	if prc.Root == "" {
		return File{}, &ProcError{Op: "open", Err: ErrEmptyRoot}
	}
	path = filepath.Join(prc.Root, path)
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
	// ugly implementation for gz files
	if filepath.Ext(path) == ".gz" {
		if stat.Size() < 1 {
			return File{}, &ProcError{Op: "open", Path: path, Err: ErrFileIsEmpty}
		}
		fz, err := gzip.NewReader(f)
		if err != nil {
			return File{}, err
		}
		s := bufio.NewScanner(fz)
		if err := s.Err(); err != nil {
			return File{}, &ProcError{Op: "scan", Path: path, Err: err}
		}
		return File{path, f, s}, nil
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
