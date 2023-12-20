// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package tty_test

import (
	"fmt"
	"os"

	"gitlab.com/nzv/prc/tty"
)

func ExampleLdiscs() {
	dcs, err := tty.Ldiscs()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	for _, v := range dcs {
		fmt.Printf("%s\n", v.Name)
	}
}
