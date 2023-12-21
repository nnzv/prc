// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package kernel

import (
	"strings"

	"gitlab.com/nzv/prc/internal"
)

func Config() (map[string]string, error) {
	f, err := internal.Open("config.gz")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data := make(map[string]string, 0)
	for f.Scan() {
		txt := f.Text()
		if len(txt) == 0 || txt[:1] == "#" {
			continue
		}
		cut := strings.Split(txt, "=")
		data[cut[0]] = strings.Trim(cut[1], `"`)
	}
	return data, nil
}
