package prc

import (
    "fmt"
    "testing"
)

func TestSwaps(t *testing.T) {
    swp, err := Swaps()
    if err != nil {
        t.Error(err)
    }
    for i, f := range swp.FileName {
        tpl := "size of %s is %d\n"
        fmt.Printf(tpl, f, swp.Size[i])
    }
}
