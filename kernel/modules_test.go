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

func TestModules(t *testing.T) {

	prc.Root = "testdata"

	tests := []struct {
		desc string
		want []Module
	}{
		{
			desc: "ok modules",
			want: []Module{
				{
					Name:    "fuse",
					Size:    188416,
					Used:    2,
					Holders: nil,
					Live:    true,
					Addr:    "0x0000000000000000",
				},
				{
					Name:    "snd_seq_dummy",
					Size:    16384,
					Used:    0,
					Holders: nil,
					Live:    true,
					Addr:    "0x0000000000000000",
				},
				{
					Name:    "snd_hrtimer",
					Size:    16384,
					Used:    1,
					Holders: nil,
					Live:    true,
					Addr:    "0x0000000000000000",
				},
				{
					Name:    "snd_seq",
					Size:    106496,
					Used:    7,
					Holders: []string{"snd_seq_dummy"},
					Live:    true,
					Addr:    "0x0000000000000000",
				},
				{
					Name:    "nf_conntrack_netlink",
					Size:    61440,
					Used:    0,
					Holders: nil,
					Live:    true,
					Addr:    "0x0000000000000000",
				},
				{
					Name:    "nfnetlink",
					Size:    24576,
					Used:    2,
					Holders: []string{"nf_conntrack_netlink"},
					Live:    true,
					Addr:    "0x0000000000000000",
				},
				{
					Name:    "xt_addrtype",
					Size:    16384,
					Used:    2,
					Holders: nil,
					Live:    true,
					Addr:    "0x0000000000000000",
				},
				{
					Name:    "br_netfilter",
					Size:    32768,
					Used:    0,
					Holders: nil,
					Live:    true,
					Addr:    "0x0000000000000000",
				},
				{
					Name:    "overlay",
					Size:    184320,
					Used:    0,
					Holders: nil,
					Live:    true,
					Addr:    "0x0000000000000000",
				},
				{
					Name:    "xt_CHECKSUM",
					Size:    16384,
					Used:    1,
					Holders: nil,
					Live:    true,
					Addr:    "0x0000000000000000",
				},
				{
					Name:    "xt_MASQUERADE",
					Size:    20480,
					Used:    13,
					Holders: nil,
					Live:    true,
					Addr:    "0x0000000000000000",
				},
				{
					Name:    "xt_conntrack",
					Size:    16384,
					Used:    11,
					Holders: nil,
					Live:    true,
					Addr:    "0x0000000000000000",
				},
				{
					Name:    "ipt_REJECT",
					Size:    16384,
					Used:    2,
					Holders: nil,
					Live:    true,
					Addr:    "0x0000000000000000",
				},
				{
					Name:    "nf_reject_ipv4",
					Size:    16384,
					Used:    1,
					Holders: []string{"ipt_REJECT"},
					Live:    true,
					Addr:    "0x00000000000",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := Modules()
			if err != nil {
				t.Fatal(err)
			}
			for i := range got {
				internal.Diff(t, fmt.Sprintf("#%d", i), &got[i], &tc.want[i])
			}
		})
	}
}
