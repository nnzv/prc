package prc

import (
    "time"
    "golang.org/x/sys/unix"
)

// Uptime returns uptime of the system (including time spent in suspend) and the
// amount of time spent in the idle process. Since age is time duration, you can
// use its type methods related.
func Uptime() (age time.Duration, err error) {
    var sys unix.Sysinfo_t
    err = unix.Sysinfo(&sys)
    if err != nil {
        return 0, err
    }
    age = time.Duration(sys.Uptime)
    return age * time.Second, nil
}
