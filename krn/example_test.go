package krn_test

import (
	"fmt"
	"os"

	"gitlab.com/nzv/prc/krn"
)

func ExampleUptime() {
	boot, idle, err := krn.Uptime()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Printf("boot=%s,idle=%s\n", boot, idle)
}

func ExamplePartitions() {
	parts, err := krn.Partitions()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	for _, v := range parts {
		fmt.Println(v.Name)
	}
}

func ExampleStats() {
	stats, err := krn.Stats()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Printf("CPU Stats: %+v\n", stats.CPU)
}

func ExampleCmdline() {
	cmdline, err := krn.Cmdline()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Printf("%+v\n", cmdline)
}

func ExampleSwaps() {
	swaps, err := krn.Swaps()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Printf("%+v\n", swaps)
}
