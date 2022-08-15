package prc

import (
    "log"
    "time"
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

func Float(str string) (u float64) {
    f, err := strconv.ParseFloat(str, 64)
    if err != nil {
        log.Fatal(err)
    }
    return f
}

func Duration(str string) time.Duration {
    return time.Duration(Float(str)) * time.Second
}

func Uint(str string) (u uint) { return uint(U64(str)) }
