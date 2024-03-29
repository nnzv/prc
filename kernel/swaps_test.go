// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package kernel

import (
	"fmt"
	"testing"

	"gitlab.com/nzv/prc"
	"gitlab.com/nzv/prc/internal"
)

func TestSwaps(t *testing.T) {

	prc.Root = "testdata"

	tests := []struct {
		desc string
		want []Swap
	}{
		{
			desc: "ok swaps",
			want: []Swap{
				{
					Filename: "/dev/sda5",
					Type:     "partition",
					Size:     2097148,
					Used:     1024,
					Priority: -2,
				},
				{
					Filename: "/dev/sdb1",
					Type:     "file",
					Size:     1048572,
					Used:     512,
					Priority: -1,
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := Swaps()
			if err != nil {
				t.Fatal(err)
			}
			for i := range got {
				internal.Diff(t, fmt.Sprintf("#%d", i), &got[i], &tc.want[i])
			}
		})
	}
}
