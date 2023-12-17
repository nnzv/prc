// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

/*
Package prc parses [procfs] files on Linux systems.

It's important to note that each go file in the  project is specifically configured to
build exclusively for Linux. In the event that the GOOS variable is set to a value
other than "linux" during interactions with the "go build" command, accessibility to the
functions in this package might be restricted due to the platform constraint.

[procfs]: https://www.kernel.org/doc/man-pages/online/pages/man5/proc.5.html
*/
package prc

// Root is the default path for the "/proc" directory. It is a basic raw string
// that represents the root path without undergoing additional checks, such as
// whether it is mounted.
var Root = "/proc"
