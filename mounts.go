package prc

import (
    "os"
    "strings"
    "bufio"
)

// MountInfo holds mounted filesystems atributtes
type MountInfo struct {
    Path  []string    // Device path
    Point []string    // Mount point
    Type  []string    // File system type
    Mode  []string    // Permission atributtes
    Dummy [][]string  // Dummy values
}

func Mount() (mnt *MountInfo, err error) {
    mnt = new(MountInfo)
    err = mnt.ReadMounts()
    if err != nil {
        return nil, err
    }
    return mnt, nil
}

func (m *MountInfo) ReadMounts() (err error) {
    f, err := os.Open("/proc/mounts")
    if err != nil {
        return err
    }
    defer f.Close()
    buf := bufio.NewScanner(f)
    for buf.Scan() {
        attr := strings.Fields(buf.Text())
        m.Path  = append(m.Path,  attr[0])
        m.Point = append(m.Point, attr[1])
        m.Type  = append(m.Type,  attr[2])
        m.Mode  = append(m.Mode,  attr[3])
        m.Dummy = append(m.Dummy, attr[4:])
    }
    return nil
}
