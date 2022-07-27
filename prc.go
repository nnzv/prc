/*
Package prc offers functions to data lookups for files stored in 
the pseudo-filesystem proc. 

Example

Here’s a trivial example that gives current uptime of the system:

    package main

    import (
        "fmt"
        "log"

        "github.com/nnzv/prc"
    )

    func main() {
        age, err := prc.Uptime()
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("Uptime is %s\n", age)
    }
*/
package prc

func Mount() (mnt *MountInfo, err error) {
    mnt = new(M)
    err = mnt.read()
    if err != nil {
        return nil, err
    }
    return mnt, nil
}
