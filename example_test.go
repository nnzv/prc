package prc_test

import (
    "fmt"
    "log"

    "git.sr.ht/~nzv/prc"
    "git.sr.ht/~nzv/prc/net"
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
    fmt.Printf("Uptime: %s\n", age.Idle)
}

func ExampleCmd() {
    args, err := prc.Cmd()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%+q\n", args)
}

func ExampleDev() {
    dev, err := net.Dev()
    if err != nil {
        log.Fatal(err)
    }
    for i, f := range dev.Face {
        tpl := "%s received %d bytes\n"
        fmt.Printf(tpl, f, dev.Receive.Bytes[i])
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

func ExamplePartitions() {
    prt, err := prc.Partitions()
    if err != nil {
        log.Fatal(err)
    }
    for i, n := range prt.Name {
        tpl := "%s has %d blocks\n"
        fmt.Printf(tpl, n, prt.Blocks[i])
    }
}
