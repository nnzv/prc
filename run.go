// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Simple mk(1) in Go. The Go team doesn't provide a standard tool for
// tasks like GNU Make, so prc minimizes dependencies, even for a single
// target system.
//
// Usage:
//
//   go run run.go <target> <flags>
//
// You can also create an alias to avoid typing each time:
//
//   alias run="go run run.go"

//go:build ignore

package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var out bytes.Buffer

type args struct {
	dir    string // env
	test   bool
	site   bool
	report bool
	fmt    bool
	check  bool
}

var cli args

func init() {
	flag.BoolVar(&cli.test, "test", false, "Run tests with verbose output and count=1")
	flag.BoolVar(&cli.site, "site", false, "Open pkgsite")
	flag.BoolVar(&cli.report, "report", false, "Run tests for generating a report")
	flag.BoolVar(&cli.fmt, "fmt", false, "Format code using gofmt with write and simplify options")
	flag.BoolVar(&cli.check, "check", false, "Check code formatting with gofmt and run go vet")
	flag.Usage = func() {
		log.Println("usage of run: ")
		flag.PrintDefaults()
	}
	cli.dir = "./..."
	val, ok := os.LookupEnv("DIR")
	if ok {
		cli.dir = val
	}
	log.SetFlags(0)
}

func main() {
	flag.Parse()
	dir, err := filepath.Abs(cli.dir)
	if err != nil {
		log.Fatal(err)
	}
	switch {
	case cli.test:
		run("go", "test", "-v", "-count=1", dir)
	case cli.site:
		run("pkgsite", "-open")
	case cli.report:
		run("go", "test", "-run=TestGenerateReport")
	case cli.fmt:
		run("gofmt", "-w", "-s", ".")
	case cli.check:
		run("gofmt", "-l", ".")
		if out.Len() > 0 {
			os.Exit(1)
		}
		run("go", "vet", dir)
	default:
		flag.Usage()
	}
}

func run(cmd string, args ...string) {
	c := exec.Command(cmd, args...)
	log.Println(c.String())
	c.Stdout = &out
	c.Stderr = c.Stdout
	if err := c.Run(); err != nil {
		log.Printf("run: %s", err)
	}
	if out.Len() > 0 {
		log.Print(c.Stdout)
	}
}
