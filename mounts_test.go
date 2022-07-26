package prc

import (
    "fmt"
    "testing"
)

func TestMounts(t *testing.T) {
    mnt, err := Mount()
    if err != nil {
        t.Error(err)
    }
    for i, m := range mnt.Path {
        fmt.Printf("%s mounted in %s\n", m, mnt.Point[i])
    }
}
