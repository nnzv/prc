package prc

import (
    "log"
    "strconv"
)

func Atoi(str string) (i int) {
    i, err := strconv.Atoi(str)
    if err != nil {
        log.Fatal(err)
    }
    return i
}

func U64(str string) (u uint64) {
    u, err := strconv.ParseUint(str, 10, 64)
    if err != nil {
        log.Fatal(err)
    }
    return u
}
