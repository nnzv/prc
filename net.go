package prc

import (
    "os"
    "strings"
    "bufio"
)

type NetInfo struct {
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

func Net() (net *NetInfo, err error) {
    net = new(NetInfo)
    err = net.read()
    if err != nil {
        return nil, err
    }
    return net, nil
}

func (n *NetInfo) read() (err error) {
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
            n.Face = append(n.Face, f)
        }
        num := make([]uint64, 0, 15)
        for _, v := range txt[1:] {
            num = append(num, U64(v))
        }

        n.Receive.Bytes = append(n.Receive.Bytes, num[0])
        n.Receive.Packets = append(n.Receive.Packets, num[1])
        n.Receive.Errors = append(n.Receive.Errors, num[2])
        n.Receive.Drop = append(n.Receive.Drop, num[3])
        n.Receive.Fifo = append(n.Receive.Fifo, num[4])
        n.Receive.Frame = append(n.Receive.Frame, num[5])
        n.Receive.Compressed = append(n.Receive.Compressed, num[6])
        n.Receive.Multicast = append(n.Receive.Multicast, num[7])

        n.Transmit.Bytes = append(n.Transmit.Bytes, num[8])
        n.Transmit.Packets = append(n.Transmit.Packets, num[9])
        n.Transmit.Errors = append(n.Transmit.Errors, num[10])
        n.Transmit.Drop = append(n.Transmit.Drop, num[11])
        n.Transmit.Fifo = append(n.Transmit.Fifo, num[12])
        n.Transmit.Colls = append(n.Transmit.Colls, num[13])
        n.Transmit.Carrier = append(n.Transmit.Carrier, num[14])
        n.Transmit.Compressed = append(n.Transmit.Compressed, num[15])
    }
    return nil
}
