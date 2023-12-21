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

func TestConfig(t *testing.T) {

	prc.Root = "testdata"

	tests := []struct {
		desc string
		want map[string]string
	}{
		{
			desc: "ok config",
			want: map[string]string{
				"CONFIG_CC_VERSION_TEXT":             "gcc (Gentoo 12.2.1_p20230121-r1 p10) 12.2.1 20230121",
				"CONFIG_CC_IS_GCC":                   "y",
				"CONFIG_GCC_VERSION":                 "120201",
				"CONFIG_CLANG_VERSION":               "0",
				"CONFIG_AS_IS_GNU":                   "y",
				"CONFIG_AS_VERSION":                  "23900",
				"CONFIG_LD_IS_BFD":                   "y",
				"CONFIG_LD_VERSION":                  "23900",
				"CONFIG_LLD_VERSION":                 "0",
				"CONFIG_CC_CAN_LINK":                 "y",
				"CONFIG_CC_CAN_LINK_STATIC":          "y",
				"CONFIG_CC_HAS_ASM_GOTO_OUTPUT":      "y",
				"CONFIG_CC_HAS_ASM_GOTO_TIED_OUTPUT": "y",
				"CONFIG_CC_HAS_ASM_INLINE":           "y",
				"CONFIG_CC_HAS_NO_PROFILE_FN_ATTR":   "y",
				"CONFIG_PAHOLE_VERSION":              "0",
				"CONFIG_IRQ_WORK":                    "y",
				"CONFIG_BUILDTIME_TABLE_SORT":        "y",
				"CONFIG_THREAD_INFO_IN_TASK":         "y",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := Config()
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("config mismatch: got %+v, want %+v", got, tc.want)
			}
		})
	}
}
