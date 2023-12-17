// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package kernel

import "gitlab.com/nzv/prc/internal"

type Swap struct {
	Filename string // Swap space name or device
	Type     string // Type: partition or file for swap
	Size     uint64 // Total size in KB
	Used     uint64 // Currently used size in KB
	Priority int64  // Priority (higher values indicate higher priority)
}

// Swaps returns swaps areas in use
func Swaps() ([]Swap, error) {
	f, err := internal.Open("swaps")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	f.Scan() // skip headers

	var data []Swap

	for f.Scan() {

		fields := f.ScanFields()

		var s Swap

		s.Filename = fields[0]
		s.Type = fields[1]

		s.Size, err = internal.ParseUint64(fields[2], 10, 64)
		if err != nil {
			return nil, &internal.ParseError{Path: f.Path, Field: "size", Err: err}
		}

		s.Used, err = internal.ParseUint64(fields[3], 10, 64)
		if err != nil {
			return nil, &internal.ParseError{Path: f.Path, Field: "used", Err: err}
		}

		s.Priority, err = internal.ParseInt(fields[4], 10, 64)
		if err != nil {
			return nil, &internal.ParseError{Path: f.Path, Field: "priority", Err: err}
		}

		data = append(data, s)
	}

	return data, nil

}
