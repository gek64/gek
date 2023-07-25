package gApp

import (
    "os"
)

type Temp struct {
    Location string
}

func NewTemp(location string) (t Temp) {
    return Temp{Location: location}
}

func (t Temp) Create() (err error) {
    _, err = os.Stat(t.Location)
    if os.IsNotExist(err) {
        err = os.MkdirAll(t.Location, 0755)
    }
    return err
}

func (t Temp) Delete() (err error) {
    _, err = os.Stat(t.Location)
    if !os.IsNotExist(err) {
        err = os.RemoveAll(t.Location)
    }
    return err
}
