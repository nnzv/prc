package prc

import (
    "os"
    "strings"
    "bufio"
)

type SwapsInfo struct {
    FileName []string
    Type     []string
    Size     []uint64
    Used     []uint64
    Priority []int
}

func Swaps() (swp *SwapsInfo, err error) {
    swp = new(SwapsInfo)
    err = swp.read()
    if err != nil {
        return nil, err
    }
    return swp, nil
}

func (s *SwapsInfo) read() (err error) {
    f, err := os.Open("/proc/swaps")
    if err != nil {
         return err
    }
    defer f.Close()
    buf := bufio.NewScanner(f)
    for buf.Scan() {
        txt := strings.Fields(buf.Text())
        if txt[0] == "Filename" {
            continue
        }
        s.FileName = append(s.FileName, txt[0])
        s.Type = append(s.Type, txt[1])
        s.Size = append(s.Size, U64(txt[2]))
        s.Used = append(s.Used, U64(txt[3]))
        s.Priority = append(s.Priority, Atoi(txt[4]))
    }
    return nil
}
