// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package krn

import (
	"reflect"
	"testing"

	"gitlab.com/nzv/prc"
)

func TestPartitions(t *testing.T) {

	prc.ProcPath = "testdata"

	tests := []struct {
		desc string
		want []Partition
	}{
		{
			desc: "ok partitions",
			want: []Partition{
				{
					Major:  7,
					Minor:  0,
					Blocks: 512000,
					Name:   "loop0",
				},
				{
					Major:  7,
					Minor:  1,
					Blocks: 523264,
					Name:   "loop1",
				},
				{
					Major:  7,
					Minor:  2,
					Blocks: 538112,
					Name:   "loop2",
				},
				{
					Major:  7,
					Minor:  3,
					Blocks: 552960,
					Name:   "loop3",
				},
				{
					Major:  7,
					Minor:  4,
					Blocks: 567808,
					Name:   "loop4",
				},
				{
					Major:  7,
					Minor:  5,
					Blocks: 582656,
					Name:   "loop5",
				},
				{
					Major:  7,
					Minor:  6,
					Blocks: 597504,
					Name:   "loop6",
				},
				{
					Major:  7,
					Minor:  7,
					Blocks: 612352,
					Name:   "loop7",
				},
				{
					Major:  179,
					Minor:  0,
					Blocks: 78125000,
					Name:   "mmcblk0",
				},
				{
					Major:  179,
					Minor:  1,
					Blocks: 1024000,
					Name:   "mmcblk0p1",
				},
				{
					Major:  179,
					Minor:  2,
					Blocks: 77004800,
					Name:   "mmcblk0p2",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := Partitions()
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("partitions mismatch: got %+v, want %+v", got, tc.want)
			}
		})
	}
}
