// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package kernel

import (
	"time"

	"gitlab.com/nzv/prc"
)

func Uptime() (time.Duration, time.Duration, error) {
	f, err := prc.Open(prc.Root, "uptime")
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()

	f.Scan()

	fields := f.ScanFields()

	boot, err := prc.ParseDuration(fields[0])
	if err != nil {
		return 0, 0, &prc.ParseError{Path: f.Path, Field: "boot", Err: err}
	}

	idle, err := prc.ParseDuration(fields[1])
	if err != nil {
		return 0, 0, &prc.ParseError{Path: f.Path, Field: "idle", Err: err}
	}

	return boot, idle, nil
}
