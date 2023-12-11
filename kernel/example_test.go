// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package kernel_test

import (
	"fmt"
	"os"

	"gitlab.com/nzv/prc/kernel"
)

func ExampleUptime() {
	boot, idle, err := kernel.Uptime()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Printf("boot=%s,idle=%s\n", boot, idle)
}

func ExamplePartitions() {
	parts, err := kernel.Partitions()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	for _, v := range parts {
		fmt.Println(v.Name)
	}
}

func ExampleStats() {
	stats, err := kernel.Stats()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Printf("CPU Stats: %+v\n", stats.CPU)
}

func ExampleCmdline() {
	cmdline, err := kernel.Cmdline()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Printf("%+v\n", cmdline)
}

func ExampleSwaps() {
	swaps, err := kernel.Swaps()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Printf("%+v\n", swaps)
}
