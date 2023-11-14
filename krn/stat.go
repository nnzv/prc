// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package krn

import (
	"strings"

	"gitlab.com/nzv/prc"
)

type Stat struct {
	CPU           map[string]CPUStat // CPU usage statistics per core
	Intr          []uint64           // Interrupt counts since boot
	ContextSwitch uint64             // Total context switches
	BootTime      uint64             // System boot time (Unix epoch)
	Processes     uint64             // Total processes/threads created
	ProcsRunning  uint64             // Runnable threads count
	ProcsBlocked  uint64             // Blocked processes count
	SoftIRQ       []uint64           // SoftIRQ counts since boot
}

type CPUStat struct {
	User      uint64 // User mode CPU time
	Nice      uint64 // Niced (priority-adjusted) user mode CPU time
	System    uint64 // Kernel mode CPU time
	Idle      uint64 // Idle CPU time
	IOWait    uint64 // I/O wait time (complex interpretation)
	IRQ       uint64 // Hardware interrupt servicing time
	SoftIRQ   uint64 // Software interrupt (softirq) servicing time
	Steal     uint64 // Involuntary wait time (e.g., virtualized)
	Guest     uint64 // Normal guest CPU time
	GuestNice uint64 // Niced guest CPU time
}

func Stats() (*Stat, error) {
	f, err := prc.Open(prc.ProcPath, "stat")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data := &Stat{
		CPU:     make(map[string]CPUStat),
		Intr:    make([]uint64, 0),
		SoftIRQ: make([]uint64, 0),
	}

	for f.Scanner.Scan() {

		var vals []uint64

		fields := strings.Fields(f.Scanner.Text())

		for _, x := range fields[1:] { // skip row columns
			v, err := prc.ParseUint64(x, 10, 64)
			if err != nil {
				return nil, &prc.ParseError{Path: f.Path, Field: fields[0], Err: err}
			}
			vals = append(vals, v)
		}

		if strings.HasPrefix(fields[0], "cpu") { // is it cpu, cpu0, cpu1, and so on?
			var c CPUStat
			for i := range vals {
				switch i {
				case 0:
					c.User = vals[0]
				case 1:
					c.Nice = vals[1]
				case 2:
					c.System = vals[2]
				case 3:
					c.Idle = vals[3]
				case 4:
					c.IOWait = vals[4]
				case 5:
					c.IRQ = vals[5]
				case 6:
					c.SoftIRQ = vals[6]
				case 7:
					c.Steal = vals[7]
				case 8:
					c.Guest = vals[8]
				case 9:
					c.GuestNice = vals[9]
				}
			}
			data.CPU[fields[0]] = c // assign to new keys like "cpu", "cpu0", etc.
		}

		switch fields[0] {
		case "intr":
			data.Intr = append(data.Intr, vals...)
		case "ctxt":
			data.ContextSwitch = vals[0]
		case "btime":
			data.BootTime = vals[0]
		case "processes":
			data.Processes = vals[0]
		case "procs_running":
			data.ProcsRunning = vals[0]
		case "procs_blocked":
			data.ProcsBlocked = vals[0]
		case "softirq":
			data.SoftIRQ = append(data.SoftIRQ, vals...)
		default:
			// skip
		}
	}

	return data, nil
}
