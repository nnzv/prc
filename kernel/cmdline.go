// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package kernel

import "gitlab.com/nzv/prc"

func Cmdline() ([]string, error) {
	f, err := prc.Open(prc.Root, "cmdline")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	f.Scan() // first line only
	return f.ScanFields(), nil
}
