// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package prc_test

import (
	"path/filepath"
	"testing"

	"gitlab.com/nzv/prc"
)

func TestOpen(t *testing.T) {

	prc.ProcPath = "testdata"

	tests := []struct {
		desc     string
		filename string
		path     string
	}{
		{
			desc:     "ok file",
			filename: "file",
			path:     filepath.Join(prc.ProcPath, "file"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			f, err := prc.Open(prc.ProcPath, tc.filename)
			if err != nil {
				t.Error(err)
			}
			if f.Path != tc.path {
				t.Errorf("mismatch proc path: got %v, want %v", f.Path, tc.path)
			}
		})
	}
}
