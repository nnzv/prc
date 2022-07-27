//go:build mage
package main

import (
    "os"
    "fmt"
    "os/exec"
)

func Doc() (err error) {
    fmt.Println("visit: http://localhost:6060/pkg/github.com/nnzv/prc")
    cmd := exec.Command("godoc")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

func Test(name string)  (err error) {
    args := []string{
                   "test",
                     "-v",
        name + "_test.go",
             name + ".go",
    }
    cmd := exec.Command("go", args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}
