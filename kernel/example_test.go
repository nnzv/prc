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

func ExampleConfig() {
	cfg, err := kernel.Config()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	val, ok := cfg["CONFIG_WERROR"]
	if !ok {
		fmt.Println("CONFIG_WERROR is not set")
		os.Exit(2)
	}
	fmt.Println(val)
}

func ExampleModules() {
	mod, err := kernel.Modules()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	for _, v := range mod {
		fmt.Printf("%s\n", v.Name)
	}
}

func ExampleMemInfo() {
	mem, err := kernel.MemInfo()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	for k, v := range mem {
		fmt.Printf("%s: %d KB\n", k, v)
	}
}

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
