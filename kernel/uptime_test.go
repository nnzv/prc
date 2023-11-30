// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kernel

import (
	"testing"
	"time"

	"gitlab.com/nzv/prc"
)

func TestUptime(t *testing.T) {

	prc.ProcPath = "testdata"

	tests := []struct {
		desc string
		boot string
		idle string
	}{
		{
			desc: "ok uptime",
			boot: "14h22m32.77s",
			idle: "47h35m36.67s",
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			boot, idle, err := Uptime()
			if err != nil {
				t.Fatal(err)
			}

			wantBoot, _ := time.ParseDuration(tc.boot)
			wantIdle, _ := time.ParseDuration(tc.idle)

			if boot != wantBoot {
				t.Errorf("uptime (boot) mismatch: got %v, want %v", boot, wantBoot)
			}

			if idle != wantIdle {
				t.Errorf("uptime (idle) mismatch: got %v, want %v", idle, wantIdle)
			}
		})
	}
}
