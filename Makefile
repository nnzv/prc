# Copyright 2023 Enzo Venturi. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

DIR=...
 
test: check
	go test -v -count=1 ./$(DIR)

site: test
	pkgsite -open

report:
	go test -v -count=1 -run=TestGenerateReport

fmt:
	gofmt -w -s .

check:
	test -z "$(shell gofmt -l .)"
	go vet ./$(DIR)
