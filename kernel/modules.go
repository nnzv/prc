// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package kernel

import "gitlab.com/nzv/prc/internal"

type Module struct {
	Name    string
	Size    uint64
	Used    uint64
	Holders []string
	Live    bool
	Addr    string
}

// Modules returns modules that have been loaded by the system
func Modules() ([]Module, error) {
	f, err := internal.Open("modules")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var data []Module

	for f.Scan() {
		fields := f.ScanFields()

		var m Module

		m.Name = fields[0]

		m.Size, err = internal.ParseUint64(fields[1], 10, 64)
		if err != nil {
			return nil, &internal.ParseError{Path: f.Path, Field: "size", Err: err}
		}

		m.Used, err = internal.ParseUint64(fields[2], 10, 64)
		if err != nil {
			return nil, &internal.ParseError{Path: f.Path, Field: "used", Err: err}
		}

		switch fields[3] {
		case "-":
			// do nothing
		default:
			m.Holders = internal.ParseCommaList(fields[3])
		}

		if fields[4] == "Live" {
			m.Live = true
		}

		m.Addr = fields[5]

		data = append(data, m)
	}
	return data, nil
}
