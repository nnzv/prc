# Copyright 2023 Enzo Venturi. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.
 
test: check
	go test -v -count=1 ./...

site: test
	pkgsite -open

report:
	go test -v -count=1 -run=TestGenerateReport

check:
	go fmt ./...; go vet ./...
