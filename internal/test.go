// Copyright 2023 Enzo Venturi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package internal

import (
	"go/token"
	"reflect"
	"testing"
)

// Diff reports differences between two values in a Go testing function. It takes four
// parameters: "t", a reference to the *testing.T object used for reporting test failures; "pfx",
// a prefix included in error messages for better identification of the comparison; "y", the
// expected value; and "x", the actual value. It specifically focuses on comparing fields of
// structs using [reflect.DeepEqual], reporting any differences, including field name, expected,
// and actual values.
func Diff(t *testing.T, pfx, y, x any) {
	t.Helper()
	yv, xv := reflect.ValueOf(y).Elem(), reflect.ValueOf(x).Elem()
	if yv.Type() != xv.Type() {
		t.Errorf("diff %s: type mismatch %v want %v", pfx, yv.Type(), xv.Type())
	}
	for i := 0; i < yv.NumField(); i++ {
		name := yv.Type().Field(i).Name
		if !token.IsExported(name) {
			continue
		}
		yf, xf := yv.Field(i).Interface(), xv.Field(i).Interface()
		if !reflect.DeepEqual(yf, xf) {
			t.Errorf("diff %s (%s): got %#v, want %#v", pfx, name, yf, xf)
		}
	}
}
