package net

import (
    "os"
    "strings"
    "bufio"

    "git.sr.ht/~nzv/prc"
)

type DevInfo struct {
        Face []string
     Receive
    Transmit
}

type Receive struct {
         Bytes  []uint64
       Packets  []uint64
        Errors  []uint64
          Drop  []uint64
          Fifo  []uint64
         Frame  []uint64
    Compressed  []uint64
     Multicast  []uint64
}

type Transmit struct {
         Bytes  []uint64
       Packets  []uint64
        Errors  []uint64
          Drop  []uint64
          Fifo  []uint64
         Colls  []uint64
        Carrier []uint64
    Compressed  []uint64
}

func Dev() (net *DevInfo, err error) {
    net = new(DevInfo)
    err = net.read()
    if err != nil {
        return nil, err
    }
    return net, nil
}

func (d *DevInfo) read() (err error) {
    f, err := os.Open("/proc/net/dev")
    if err != nil {
        return err
    }
    defer f.Close()
    buf := bufio.NewScanner(f)
    for buf.Scan() {
        txt := strings.Fields(buf.Text())
        switch t := txt[0]; t {
        case "Inter-|", "face":
            continue
        default:
            f := strings.TrimSuffix(t, ":")
            d.Face = append(d.Face, f)
        }
        num := make([]uint64, 0, 15)
        for _, v := range txt[1:] {
            num = append(num, prc.U64(v))
        }

        d.Receive.Bytes = append(d.Receive.Bytes, num[0])
        d.Receive.Packets = append(d.Receive.Packets, num[1])
        d.Receive.Errors = append(d.Receive.Errors, num[2])
        d.Receive.Drop = append(d.Receive.Drop, num[3])
        d.Receive.Fifo = append(d.Receive.Fifo, num[4])
        d.Receive.Frame = append(d.Receive.Frame, num[5])
        d.Receive.Compressed = append(d.Receive.Compressed, num[6])
        d.Receive.Multicast = append(d.Receive.Multicast, num[7])

        d.Transmit.Bytes = append(d.Transmit.Bytes, num[8])
        d.Transmit.Packets = append(d.Transmit.Packets, num[9])
        d.Transmit.Errors = append(d.Transmit.Errors, num[10])
        d.Transmit.Drop = append(d.Transmit.Drop, num[11])
        d.Transmit.Fifo = append(d.Transmit.Fifo, num[12])
        d.Transmit.Colls = append(d.Transmit.Colls, num[13])
        d.Transmit.Carrier = append(d.Transmit.Carrier, num[14])
        d.Transmit.Compressed = append(d.Transmit.Compressed, num[15])
    }
    return nil
}
