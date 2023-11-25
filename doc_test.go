// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package prc

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
	"text/tabwriter"
)

// TestGenerateReport generates a test report by comparing the contents of directories
// based on the specified expectations in the knowDirs variable.
func TestGenerateReport(t *testing.T) {
	report := new(bytes.Buffer)

	w := tabwriter.NewWriter(report, 0, 4, 0, ' ', 0)

	fmt.Fprintln(w, "\tWANT \tGOT \tDID")

	var aw, ag int // Initialize total counts.

	for dir, want := range knowDirs {
		entry, err := os.ReadDir(dir)
		if err != nil {
			t.Fatal(err)
		}
		var got int
		for _, e := range entry {
			if !strings.HasSuffix(e.Name(), "_test.go") || !e.IsDir() {
				for _, f := range want {
					if e.Name() == f+".go" {
						got++
					}
				}
			}
		}
		cnt := len(want) // Total expected files in the directory.
		fmt.Fprintf(w, "%s \t%d\t%d\t%.2f%%\n", dir, cnt, got, pct(got, cnt))
		aw, ag = cnt+aw, got+ag
	}
	fmt.Fprintf(w, "\t \t\t\t-----\n\t \t\t\t%.2f%%\n", pct(ag, aw))
	w.Flush()
	t.Logf("\n%s", report)
}

// pct calculates the percentage of x in relation to y. It returns the result as a float64.
func pct(x, y int) float64 { return (float64(x) * float64(y)) / 100.0 }

// knowDirs represents the current directories mirroring the proc filesystem. It functions as a reference
// for tracking the project's objectives and generating a report on the current state of development. Each
// key represents a directory within the project, and its corresponding value is a slice of files expected
// to be present in those directories.
//
// The filenames in the slice correspond to those specified in the official procfs documentation. In cases
// where the documentation is absent or ambiguous regarding a specific file, a default main file will be used.
var knowDirs = map[string][]string{
	// Ext4 File System Parameters
	"ext4": []string{
		"mb_groups", // Details of multiblock allocator buddy cache of free blocks
	},
	// IDE Devices
	"ide": []string{
		"main",
	},
	// Kernel data and statistics
	"krn": []string{
		"apm",          // Advanced power management
		"bootconfig",   // Kernel command line and bootloader parameters
		"buddyinfo",    // Kernel memory allocator information
		"bus",          // Bus-specific information
		"cmdline",      // Kernel command line
		"cpuinfo",      // CPU information
		"devices",      // Available devices
		"consoles",     // Shows registered system console lines.
		"dma",          // Used DMS channels
		"filesystems",  // Supported filesystems
		"driver",       // Various drivers, currently rtc
		"execdomains",  // Execdomains related to security
		"fb",           // Frame Buffer devices
		"fs",           // File system parameters, currently nfs/exports
		"ide",          // Info about the IDE subsystem
		"interrupts",   // Interrupt usage
		"iomem",        // Memory map
		"ioports",      // I/O port usage
		"irq",          // Masks for irq to CPU affinity
		"isapnp",       // ISA PnP (Plug&Play)
		"kcore",        // Kernel core image
		"kmsg",         // Kernel messages
		"ksyms",        // Kernel symbol table
		"loadavg",      // Load average and process statistics
		"locks",        // Kernel locks
		"meminfo",      // Memory information
		"misc",         // Miscellaneous
		"modules",      // List of loaded modules
		"mounts",       // Mounted filesystems
		"pagetypeinfo", // Page allocator information
		"partitions",   // Table of partitions
		// "pci",       // Deprecated PCI bus info (new way -> /proc/bus/pci/, decoupled by lspci)
		"rtc",         // Real-time clock
		"slabinfo",    // Slab pool info
		"softirqs",    // Softirq usage
		"stat",        // Overall statistics
		"swaps",       // Swap space utilization
		"sys",         // See chapter 2
		"sysvipc",     // SysVIPC Resources info (msg, sem, shm)
		"uptime",      // Wall clock since boot, combined idle time of all CPUs
		"version",     // Kernel version
		"video",       // bttv info of video resources
		"vmallocinfo", // Show vmalloced areas
	},
	// Networking info
	"net": []string{
		"arp",           // Kernel ARP table
		"dev",           // Network devices with statistics
		"dev_mcast",     // Layer2 multicast groups a device is listening
		"dev_stat",      // Network device status
		"ip_fwchains",   // Firewall chain linkage
		"ip_fwnames",    // Firewall chain names
		"ip_masq",       // Directory containing masquerading tables
		"ip_masquerade", // Major masquerading table
		"netstat",       // Network statistics
		"raw",           // Raw device statistics
		"route",         // Kernel routing table
		"rpc",           // Directory containing rpc info
		"rt_cache",      // Routing cache
		"snmp",          // SNMP data
		"sockstat",      // Socket statistics
		"softnet_stat",  // Per-CPU incoming packet queues statistics of online CPUs
		"tcp",           // TCP sockets
		"udp",           // UDP sockets
		"unix",          // UNIX domain sockets
		"wireless",      // Wireless interface data (Wavelan etc)
		"igmp",          // IP multicast addresses this host joined
		"psched",        // Global packet scheduler parameters.
		"netlink",       // List of PF_NETLINK sockets
		"ip_mr_vifs",    // List of multicast virtual interfaces
		"ip_mr_cache",   // List of multicast routing cache
		"udp6",          // UDP sockets (IPv6)
		"tcp6",          // TCP sockets (IPv6)
		"raw6",          // Raw device statistics (IPv6)
		"igmp6",         // IP multicast addresses this host joined (IPv6)
		"if_inet6",      // List of IPv6 interface addresses
		"ipv6_route",    // Kernel routing table for IPv6
		"rt6_stats",     // Global IPv6 routing tables statistics
		"sockstat6",     // Socket statistics (IPv6)
		"snmp6",         // SNMP data (IPv6)
	},
	// Parallel port info
	"parport": []string{
		"autoprobe", // IEEE-1284 device ID information
		"devices",   // List of device drivers using the port. '+' indicates the current user
		"hardware",  // Parallel port's base address, IRQ line, and DMA channel
		"irq",       // IRQ that parport is using for that port. Can be modified by writing a new value
	},
	// Per-Process parameters
	"ps": []string{
		"oom_adj",         // Adjust the oom-killer score
		"oom_score_adj",   // Adjust the oom-killer score
		"oom_score",       // Display current oom-killer score
		"io",              // Display the IO accounting fields
		"coredump_filter", // Core dump filtering settings
		"mountinfo",       // Information about mounts
		"comm",            // Command name of the process
		"children",        // Information about task children
		"fdinfo/<fd>",     // Information about opened file
		"map_files",       // Information about memory mapped files
		"timerslack_ns",   // Task timerslack value
		"patch_state",     // Livepatch patch operation state
		"arch_status",     // Task architecture specific information
		"fd",              // List of symlinks to open files
	},
	// Process-Specific subdirectories
	"psd": []string{
		"clear_refs",   // Clears page referenced bits shown in smaps output
		"cmdline",      // Command line arguments
		"cpu",          // Current and last CPU in which it was executed
		"cwd",          // Link to the current working directory
		"environ",      // Values of environment variables
		"exe",          // Link to the executable of this process
		"fd",           // Directory containing all file descriptors
		"maps",         // Memory maps to executables and library files
		"mem",          // Memory held by this process
		"root",         // Link to the root directory of this process
		"stat",         // Process status
		"statm",        // Process memory status information
		"status",       // Process status in human-readable form
		"wchan",        // Kernel function symbol the task is blocked in, or "0" if not blocked (CONFIG_KALLSYMS=y)
		"pagemap",      // Page table
		"stack",        // Report full stack trace, enable via CONFIG_STACKTRACE
		"smaps",        // An extension based on maps, showing the memory consumption of each mapping and associated flags (CONFIG_KALLSYMS=y)
		"smaps_rollup", // Accumulated smaps stats for all mappings of the process. Can be derived from smaps but faster and more convenient
		"numa_maps",    // An extension based on maps, showing the memory locality
	},
	// SCSI info
	"scsi": []string{
		"main",
	},
	// Modifying system parameters
	"sys": []string{
		"main",
	},
	// TTY info
	"tty": []string{
		"drivers",       // List of drivers and their usage
		"ldiscs",        // Registered line disciplines
		"driver_serial", // Usage statistics and status of single tty lines
	},
}
