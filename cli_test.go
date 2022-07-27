package prc

import (
    "fmt"
    "testing"
)

func TestCli(t *testing.T) {
    args, err := Cli()
    if err != nil {
        t.Error(err)
    }
    fmt.Printf("%+q\n", args)
}
