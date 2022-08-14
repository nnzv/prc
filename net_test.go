package prc

import (
    "fmt"
    "testing"
)

func TestNet(t *testing.T) {
    net, err := Net()
    if err != nil {
        t.Error(err)
    }
    for i, f := range net.Face {
        tpl := "%s received %d bytes\n"
        fmt.Printf(tpl, f, net.Receive.Bytes[i])
    }
}
