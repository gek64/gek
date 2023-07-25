package gApp

import (
    "fmt"
    "os"
    "path/filepath"
)

type Data struct {
    // 数据文件
    Files []string
    // 数据文件夹
    Location string
    // 删除时是否删除数据文件夹
    UninstallDeleteLocationFolder bool
}

// NewData 新建数据
func NewData(files []string, location string, uninstallDeleteLocationFolder bool) (d Data) {
    return Data{Files: files, Location: location, UninstallDeleteLocationFolder: uninstallDeleteLocationFolder}
}

// Install 安装数据文件夹
func (d Data) Install() (err error) {
    // 检测数据文件夹路径是否存在
    // 不存在则创建
    _, err = os.Stat(d.Location)
    if os.IsNotExist(err) {
        err = os.MkdirAll(d.Location, 0755)
        if err != nil {
            return err
        }
    }
    return err
}

// Uninstall 卸载数据文件,并删除数据文件夹
func (d Data) Uninstall() (err error) {
    // 检测数据文件夹路径是否存在
    _, err = os.Stat(d.Location)
    if os.IsNotExist(err) {
        return fmt.Errorf("can't find data location %s", d.Location)
    }

    if d.UninstallDeleteLocationFolder {
        err = os.RemoveAll(d.Location)
    } else {
        for _, file := range d.Files {
            err = os.RemoveAll(filepath.Join(d.Location, file))
        }
    }
    return err
}
