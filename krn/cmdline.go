// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package krn

import (
	"strings"

	"gitlab.com/nzv/prc"
)

func Cmdline() ([]string, error) {
	f, err := prc.Open(prc.ProcPath, "cmdline")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	f.Scanner.Scan() // first line only
	return strings.Fields(f.Scanner.Text()), nil
}
