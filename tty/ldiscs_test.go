// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package tty

import (
	"reflect"
	"testing"

	"gitlab.com/nzv/prc"
)

func TestLdiscs(t *testing.T) {

	prc.Root = "testdata"

	tests := []struct {
		desc string
		want []Ldisc
	}{
		{
			desc: "ok ldiscs",
			want: []Ldisc{
				{
					Name:   "n_tty",
					Number: 0,
				},
				{
					Name:   "input",
					Number: 2,
				},
				{
					Name:   "n_null",
					Number: 27,
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := Ldiscs()
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ldiscs mismatch: got %v, want %v", got, tc.want)
			}
		})
	}
}
