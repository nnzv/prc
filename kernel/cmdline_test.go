// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kernel

import (
	"reflect"
	"testing"

	"gitlab.com/nzv/prc"
)

func TestCmdline(t *testing.T) {

	prc.Root = "testdata"

	tests := []struct {
		desc string
		want []string
	}{
		{
			desc: "ok cmdline",
			want: []string{"root=/dev/sda1", "quiet", "splash"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := Cmdline()
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("cmdline mismatch: got %v, want %v", got, tc.want)
			}
		})
	}
}
