package prc

import (
    "strings"
    "io/ioutil"
)

// Cmd returns the command-line arguments passed to 
// the Linux kernel at boot time.
func Cmd() (args []string, err error) {
    read, err := ioutil.ReadFile("/proc/cmdline")
    if err != nil {
        return nil, err
    }
    return strings.Fields(string(read)), nil
}
