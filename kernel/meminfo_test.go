// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package kernel

import (
	"reflect"
	"testing"

	"gitlab.com/nzv/prc"
)

func TestMemInfo(t *testing.T) {

	prc.Root = "testdata"

	tests := []struct {
		desc string
		want map[string]uint64
	}{
		{
			desc: "ok meminfo",
			want: map[string]uint64{
				"MemTotal":          32808612,
				"MemFree":           22483524,
				"MemAvailable":      27254992,
				"Buffers":           283192,
				"Cached":            4890380,
				"SwapCached":        0,
				"Active":            1656072,
				"Inactive":          7422108,
				"Active(anon)":      3072,
				"Inactive(anon)":    4018460,
				"Active(file)":      1653000,
				"Inactive(file)":    3403648,
				"Unevictable":       72,
				"Mlocked":           72,
				"SwapTotal":         0,
				"SwapFree":          0,
				"Zswap":             0,
				"Zswapped":          0,
				"Dirty":             628,
				"Writeback":         0,
				"AnonPages":         3901592,
				"Mapped":            470648,
				"Shmem":             116924,
				"KReclaimable":      163872,
				"Slab":              245640,
				"SReclaimable":      163872,
				"SUnreclaim":        81768,
				"KernelStack":       17808,
				"PageTables":        48004,
				"SecPageTables":     0,
				"NFS_Unstable":      0,
				"Bounce":            0,
				"WritebackTmp":      0,
				"CommitLimit":       16404304,
				"Committed_AS":      9383080,
				"VmallocTotal":      34359738367,
				"VmallocUsed":       47008,
				"VmallocChunk":      0,
				"Percpu":            2112,
				"HardwareCorrupted": 0,
				"AnonHugePages":     24576,
				"ShmemHugePages":    0,
				"ShmemPmdMapped":    0,
				"FileHugePages":     0,
				"FilePmdMapped":     0,
				"CmaTotal":          0,
				"CmaFree":           0,
				"HugePages_Total":   0,
				"HugePages_Free":    0,
				"HugePages_Rsvd":    0,
				"HugePages_Surp":    0,
				"Hugepagesize":      2048,
				"Hugetlb":           0,
				"DirectMap4k":       162460,
				"DirectMap2M":       6070272,
				"DirectMap1G":       28311552,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := MemInfo()
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("meminfo mismatch: got %+v, want %+v", got, tc.want)
			}
		})
	}
}
