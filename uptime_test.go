package prc

import (
    "fmt"
    "testing"
)

func TestUptime(t *testing.T) {
    age, err := Uptime()
    if err != nil {
        t.Error(err)
    }
    fmt.Printf("Uptime: %s\n", age)
}
