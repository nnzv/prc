// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package kernel

import "gitlab.com/nzv/prc/internal"

func MemInfo() (map[string]uint64, error) {
	f, err := internal.Open("meminfo")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data := make(map[string]uint64, 0)

	for f.Scan() {
		fields := f.ScanFields()

		name := fields[0][:len(fields[0])-1]

		data[name], err = internal.ParseUint64(fields[1], 10, 64)
		if err != nil {
			return nil, &internal.ParseError{Path: f.Path, Field: "size", Err: err}
		}
	}
	return data, nil
}
