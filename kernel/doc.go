// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

/*
Package kernel provides kernel info functions.

Every function in this package is designed according to the files listed in Table 1-5 under
https://docs.kernel.org/filesystems/proc.html#kernel-data. Some functions might not work on
your system. It depends on how your kernel is set up and which modules are loaded. Certain
files could be there, but others might be missing, affecting how some functions behave.
*/
package kernel
