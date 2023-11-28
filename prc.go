// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package prc

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Default paths for the "/proc" and "/proc/sys" directories. These paths are simple raw strings,
// and no checks, such as whether they are mounted, are performed on them.
var (
	ProcPath = "/proc"
	SysPath  = "/proc/sys"
)

// File represents an already open file, ready for scanning its content.
type File struct {
	Path           string // Absolute path (mostly used for errors)
	*os.File              // File handle
	*bufio.Scanner        // Scanner for reading the file
}

// Open combines root and path to open a file, typically using [ProcPath] or [SysPath] as root.
// It checks for errors and validates file properties, returning a [File] struct with the file path
// and a [bufio.Scanner] for reading its content upon success.
func Open(root, path string) (*File, error) {
	p := filepath.Join(root, path)
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	stat, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("proc %q", err) // [fs.PathError] includes the path information.
	}
	if stat.IsDir() {
		return nil, fmt.Errorf("proc %s: path is a directory", p)
	}
	// BUG(nzv): Proc files don't follow the regular file rules and don't have a size.
	// The usual size check using stat.Size() doesn't work, excluding some files.
	// Check GitLab issue #1 for more info.
	// if stat.Size() < 1 {
	//    return nil, fmt.Errorf("proc %s: file is empty", p)
	// }
	sc := bufio.NewScanner(f)
	if err := sc.Err(); err != nil {
		f.Close()
		return nil, fmt.Errorf("proc %s: %q", p, err)
	}
	return &File{p, f, sc}, nil
}

// Close closes the /proc file by closing its file handle.
func (f *File) Close() error { return f.File.Close() }

// Scan scans the /proc file by uing its [bufio.Scanner] handle.
func (f *File) Scan() bool { return f.Scanner.Scan() }

// SplitWords configures the scanner to split words using [bufio.ScanWords].
func (f *File) SplitWords() { f.Scanner.Split(bufio.ScanWords) }

// ScanFields scans the current line of the file, splitting it into fields using whitespace.
// Returns a slice of strings representing the fields.
func (f *File) ScanFields() []string { return strings.Fields(f.Scanner.Text()) }

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
