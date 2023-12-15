// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package prc_test

import (
	"testing"

	"gitlab.com/nzv/prc"
)

func TestOpen(t *testing.T) {

	tests := []struct {
		desc     string
		filename string
		path     string
		root     string
		err      string
	}{
		{
			desc:     "ok file",
			filename: "file",
			path:     "testdata/file",
			root:     "testdata",
			err:      "",
		},
		{
			desc:     "nok file (empty)",
			filename: "empty",
			path:     "",
			root:     "testdata",
			err:      "proc open testdata/empty: file is empty",
		},
		{
			desc:     "nok path (type dir)",
			filename: "dir",
			path:     "",
			root:     "testdata",
			err:      "proc open testdata/dir: path is a directory",
		},
		{
			desc:     "nok root (empty root)",
			filename: "file",
			path:     "",
			root:     "",
			err:      "proc open: empty root path",
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			prc.Root = tc.root
			f, err := prc.Open(tc.filename)
			if err != nil && err.Error() != tc.err {
				t.Fatal(err)
			}
			if f.Path != tc.path {
				t.Errorf("mismatch proc path: got %v, want %v", f.Path, tc.path)
			}
		})
	}
}
