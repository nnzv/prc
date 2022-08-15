package net

import (
    "fmt"
    "testing"
)

func TestDev(t *testing.T) {
    dev, err := Dev()
    if err != nil {
        t.Error(err)
    }
    for i, f := range dev.Face {
        tpl := "%s received %d bytes\n"
        fmt.Printf(tpl, f, dev.Receive.Bytes[i])
    }
}
