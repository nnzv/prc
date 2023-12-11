// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kernel

import "gitlab.com/nzv/prc"

type Partition struct {
	Major  uint64 // Major number of the partition.
	Minor  uint64 // Minor number of the partition.
	Blocks uint64 // Total blocks in the partition.
	Name   string // Name of the partition.
}

func Partitions() ([]Partition, error) {
	f, err := prc.Open(prc.Root, "partitions")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	f.Scan() // skip headers

	var data []Partition

	for f.Scan() {

		fields := f.ScanFields()

		if len(fields) != 4 {
			continue // skip invalid fields
		}

		var p Partition

		p.Major, err = prc.ParseUint64(fields[0], 10, 64)
		if err != nil {
			return nil, &prc.ParseError{Path: f.Path, Field: "major", Err: err}
		}

		p.Minor, err = prc.ParseUint64(fields[1], 10, 64)
		if err != nil {
			return nil, &prc.ParseError{Path: f.Path, Field: "minor", Err: err}
		}

		p.Blocks, err = prc.ParseUint64(fields[2], 10, 64)
		if err != nil {
			return nil, &prc.ParseError{Path: f.Path, Field: "blocks", Err: err}
		}

		p.Name = fields[3]

		data = append(data, p)
	}

	return data, nil
}
