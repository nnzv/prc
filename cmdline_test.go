package prc

import (
    "fmt"
    "testing"
)

func TestCmd(t *testing.T) {
    args, err := Cmd()
    if err != nil {
        t.Error(err)
    }
    fmt.Printf("%+q\n", args)
}
