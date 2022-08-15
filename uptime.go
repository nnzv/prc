package prc

import (
    "time"
    "strings"
    "io/ioutil"
)

type UpInfo struct {
      Up time.Duration 
    Idle time.Duration
}

func Uptime() (age *UpInfo, err error) {
    age = new(UpInfo)
    err = age.read()
    if err != nil {
        return nil, err
    }
    return age, nil
}

func (u *UpInfo) read() (err error) {
    f, err := ioutil.ReadFile("/proc/uptime")
    if err != nil {
        return err
    }
    age := strings.Fields(string(f))
    u.Up = Duration(age[0]) 
    u.Idle = Duration(age[1]) 
    return nil
}
