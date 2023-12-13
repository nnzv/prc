// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package tty

import "gitlab.com/nzv/prc"

type Ldisc struct {
	Name   string // Ldisc name
	Number uint64 // Reserved number for the ldisc
}

func Ldiscs() ([]Ldisc, error) {
	f, err := prc.Open("tty/ldiscs")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var data []Ldisc

	for f.Scan() {

		fields := f.ScanFields()

		var l Ldisc

		l.Name = fields[0]

		l.Number, err = prc.ParseUint64(fields[1], 10, 64)
		if err != nil {
			return nil, &prc.ParseError{Path: f.Path, Field: "number", Err: err}
		}

		data = append(data, l)
	}
	return data, nil
}
