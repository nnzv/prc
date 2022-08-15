package prc

import (
    "os"
    "strings"
    "bufio"
)

type PartInfo struct {
      Name []string
     Major []uint
     Minor []uint
    Blocks []uint64
}

func Partitions() (prt *PartInfo, err error) {
    prt = new(PartInfo)
    err = prt.read()
    if err != nil {
        return nil, err
    }
    return prt, nil
}

func (p *PartInfo) read() (err error) {
    f, err := os.Open("/proc/partitions")
    if err != nil {
        return err
    }
    defer f.Close()
    buf := bufio.NewScanner(f)
    for buf.Scan() {
        txt := strings.Fields(buf.Text())
        if len(txt) == 0 || txt[0] == "major" {
            continue
        }
        p.Major = append(p.Major, Uint(txt[0]))
        p.Minor = append(p.Minor, Uint(txt[1]))
        p.Blocks = append(p.Blocks, U64(txt[2]))
        p.Name = append(p.Name, txt[3])
    }
    return nil
}
