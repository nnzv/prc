package prc

func Mount() (mnt *MountInfo, err error) {
    mnt = new(MountInfo)
    err = mnt.read()
    if err != nil {
        return nil, err
    }
    return mnt, nil
}
