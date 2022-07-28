```txt
 ___ ___ ___ 
| . |  _|  _|
|  _|_| |___| proc reader
|_|
```

[![Go Reference](https://pkg.go.dev/badge/github.com/nnzv/prc.svg)](https://pkg.go.dev/github.com/nnzv/prc)

Offers data lookups for files stored on `/proc`. The easiest way to install is to run:

```sh
% go get -u github.com/nnzv/prc
```

### Example

Here's a trivial example that gives current uptime of the system:

```go
package main

import (
    "fmt"
    "log"

    "github.com/nnzv/prc"
)

func main() {
    age, err := prc.Uptime()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Uptime is %s\n", age)
    // Output: Uptime is 2h25m26s
}
```

### Contributing

Please do. Contributing is one of the most important things to enhance `prc`. So feel free to report or edit related project stuff, but be respectful. Note that we use Go 1.18 and [mage](mage) as automating helper.

Testing, for example, can be reduced to:

```txt
% mage test uptime
=== RUN   TestUptime
Uptime: 2h25m26s
--- PASS: TestUptime (0.00s)
PASS
ok  	command-line-arguments	0.001s
```

[mage]: https://magefile.org
