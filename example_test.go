package prc_test

import (
    "fmt"
    "log"

    "git.sr.ht/~nzv/prc"
)

func ExampleMount() {
    mnt, err := prc.Mount()
    if err != nil {
        log.Fatal(err)
    }
    // Iterate over device paths
    for i, p := range mnt.Path {
        // Access device mount point by current index
        fmt.Printf("%s mounted in %s\n", p, mnt.Point[i])
    }
}

func ExampleUptime() {
    age, err := prc.Uptime()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Uptime: %1.f\n", age.Hours())
}

func ExampleCmd() {
    args, err := prc.Cmd()
    if err != nil {
        t.Error(err)
    }
    fmt.Printf("%+q\n", args)
}

func ExampleNet() {
    net, err := prc.Net()
    if err != nil {
        log.Fatal(err)
    }
    for i, f := range net.Face {
        tpl := "%s received %d bytes\n"
        fmt.Printf(tpl, f, net.Receive.Bytes[i])
    }
}
