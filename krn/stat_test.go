// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package krn

import (
	"reflect"
	"testing"

	"gitlab.com/nzv/prc"
)

func TestStat(t *testing.T) {

	prc.ProcPath = "testdata"

	tests := []struct {
		desc string
		want *Stat
	}{
		{
			desc: "ok stat",
			want: &Stat{
				CPU: map[string]CPUStat{
					"cpu": {
						User:      123,
						Nice:      456,
						System:    789,
						Idle:      1011,
						IOWait:    1213,
						IRQ:       1415,
						SoftIRQ:   1617,
						Steal:     1819,
						Guest:     0,
						GuestNice: 0,
					},
					"cpu0": {
						User:      23,
						Nice:      45,
						System:    67,
						Idle:      89,
						IOWait:    101,
						IRQ:       112,
						SoftIRQ:   131,
						Steal:     141,
						Guest:     0,
						GuestNice: 0,
					},
					"cpu1": {
						User:      100,
						Nice:      200,
						System:    300,
						Idle:      400,
						IOWait:    500,
						IRQ:       600,
						SoftIRQ:   700,
						Steal:     800,
						Guest:     0,
						GuestNice: 0,
					},
				},
				Intr:          []uint64{12345},
				ContextSwitch: 54321,
				BootTime:      1631234567,
				Processes:     789,
				ProcsRunning:  2,
				ProcsBlocked:  1,
				SoftIRQ:       []uint64{1111, 2222, 3333, 4444, 5555, 6666, 7777, 8888},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := Stats()
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("stat mismatch: got %+v, want %+v", got, tc.want)
			}
		})
	}
}
