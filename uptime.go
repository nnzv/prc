package prc

import (
    "time"
    "golang.org/x/sys/unix"
)

func Uptime() (age time.Duration, err error) {
    var sys unix.Sysinfo_t
    err = unix.Sysinfo(&sys)
    if err != nil {
        return 0, err
    }
    age = time.Duration(sys.Uptime)
    return age * time.Second, nil
}
