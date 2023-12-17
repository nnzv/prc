// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package kernel

import (
	"time"

	"gitlab.com/nzv/prc/internal"
)

// Uptime returns system uptime and idle time.
func Uptime() (time.Duration, time.Duration, error) {
	f, err := internal.Open("uptime")
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()

	f.Scan()

	fields := f.ScanFields()

	boot, err := internal.ParseDuration(fields[0])
	if err != nil {
		return 0, 0, &internal.ParseError{Path: f.Path, Field: "boot", Err: err}
	}

	idle, err := internal.ParseDuration(fields[1])
	if err != nil {
		return 0, 0, &internal.ParseError{Path: f.Path, Field: "idle", Err: err}
	}

	return boot, idle, nil
}
