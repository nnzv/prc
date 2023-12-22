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

func TestConsoles(t *testing.T) {

	prc.Root = "testdata"

	tests := []struct {
		desc string
		want []Console
	}{
		{
			desc: "ok consoles",
			want: []Console{
				{
					Name:  "tty0",
					Ops:   []rune{87, 85},
					Flags: []rune{40, 69, 67, 112, 41},
					Major: 4,
					Minor: 7,
				},
				{
					Name:  "ttyS0",
					Ops:   []rune{87},
					Flags: []rune{40, 69, 112, 41},
					Major: 4,
					Minor: 64,
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := Consoles()
			if err != nil {
				t.Fatal(err)
			}
			for i := range got {
				internal.Diff(t, fmt.Sprintf("#%d", i), &got[i], &tc.want[i])
			}
		})
	}
}
