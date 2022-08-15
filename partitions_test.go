package prc

import (
    "fmt"
    "testing"
)

func TestPartitions(t *testing.T) {
    prt, err := Partitions()
    if err != nil {
        t.Error(err)
    }
    for i, n := range prt.Name {
        tpl := "%s has %d blocks\n"
        fmt.Printf(tpl, n, prt.Blocks[i])
    }
}
