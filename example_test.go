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
    for i, p := range mnt.Path {
        tpl := "%s mounted in %s\n"
        fmt.Printf(tpl, p, mnt.Point[i])
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
        log.Fatal(err)
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

func ExampleSwaps() {
    swp, err := prc.Swaps()
    if err != nil {
        log.Fatal(err)
    }
    for i, f := range swp.FileName {
        tpl := "size of %s is %d\n"
        fmt.Printf(tpl, f, swp.Size[i])
    }
}
