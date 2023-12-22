// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package kernel

import (
	"strings"

	"gitlab.com/nzv/prc/internal"
)

type Console struct {
	Name  string
	Ops   []rune
	Flags []rune
	Major uint64
	Minor uint64
}

func Consoles() ([]Console, error) {
	f, err := internal.Open("consoles")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var data []Console

	for f.Scan() {

		var parenthese bool

		var trim strings.Builder

		// remove '-' in ops, spaces between flags
		for _, c := range f.Text() {
			switch {
			case c == '(':
				parenthese = true
			case c == ')':
				parenthese = false // parenthese's end
			case c == '-' || parenthese && c == ' ':
				continue
			default:
				// do nothing
			}
			trim.WriteRune(c)
		}

		fields := strings.Fields(trim.String())

		var c Console

		c.Name = fields[0]

		c.Ops = []rune(fields[1])

		c.Flags = []rune(fields[2])

		cut := strings.Split(fields[3], ":")

		c.Major, err = internal.ParseUint64(cut[0], 10, 64)
		if err != nil {
			return nil, &internal.ParseError{Path: f.Path, Field: "major", Err: err}
		}

		c.Minor, err = internal.ParseUint64(cut[1], 10, 64)
		if err != nil {
			return nil, &internal.ParseError{Path: f.Path, Field: "minor", Err: err}
		}

		data = append(data, c)
	}
	return data, nil
}
