// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Simple mk(1) in Go. The Go team doesn't provide a standard tool for
// tasks like GNU Make, so prc minimizes dependencies, even for a single
// target system.
//
// Usage:
//
//   go run mk.go <target> <flags>
//
// You can also create an alias to avoid typing each time:
//
//   alias mk="go run mk.go"

//go:build ignore

package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var out bytes.Buffer
var dir string // flags

func init() {
	flag.StringVar(&dir, "dir", "./...", "specify the directory for Go commands")
	flag.Usage = func() {
		log.Println("usage of mk: ")
		flag.PrintDefaults()
	}
	log.SetFlags(0)
}

func main() {
	flag.Parse()
	dir, err := filepath.Abs(dir)
	if err != nil {
		log.Fatal(err)
	}
	t := strings.TrimSpace(flag.Arg(0))
	switch t {
	case "test", "":
		run("go", "test", "-v", "-count=1", dir)
	case "site":
		run("pkgsite", "-open")
	case "report":
		run("go", "test", "-run=TestGenerateReport")
	case "fmt":
		run("gofmt", "-w", "-s", ".")
	case "check":
		run("gofmt", "-l", ".")
		if out.Len() > 0 {
			os.Exit(1)
		}
		run("go", "vet", dir)
	default:
		log.Fatalf("mk: unknown target %#v\n", t)
	}
}

func run(cmd string, args ...string) {
	c := exec.Command(cmd, args...)
	log.Println(c.String())
	c.Stdout = &out
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		log.Fatalf("mk: %s", err)
	}
	if out.Len() > 0 {
		log.Print(&out)
	}
}
