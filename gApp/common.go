package gApp

import (
    "github.com/gek64/gek/gExec"
    "os"
    "os/exec"
)

var (
    // SupportedOS 支持的系统
    SupportedOS = []string{"linux", "freebsd", "windows", "darwin", "android", "ios", "openbsd", "netbsd", "dragonfly", "solaris", "aix", "illumos", "js", "plan9"}

    // SupportedArch 支持的架构
    SupportedArch = []string{"amd64", "386", "arm64", "arm", "riscv64", "mips64", "mips64le", "mips", "mipsle", "ppc64", "ppc64le", "s390x", "wasm"}

    // InitSystem 初始化系统
    InitSystem = map[string]string{
        "systemd": "systemctl",
        "openrc":  "openrc",
        "rc.d":    "rcorder",
    }
)

// ExtractZip 压缩zip文件
func ExtractZip(archiveFile string, location string) (err error) {
    // 如果输出路径不存在则创建
    _, err = os.Stat(location)
    if os.IsNotExist(err) {
        err = os.MkdirAll(location, 0755)
        if err != nil {
            return err
        }
    }

    // unzip 解压
    err = gExec.Run(exec.Command("unzip", "-o", "-d", location, archiveFile))
    if err != nil {
        return err
    }

    return nil
}

// ExtractTar 压缩tar文件
func ExtractTar(archiveFile string, location string) (err error) {
    // 如果输出路径不存在则创建
    _, err = os.Stat(location)
    if os.IsNotExist(err) {
        err = os.MkdirAll(location, 0755)
        if err != nil {
            return err
        }
    }

    // tar 解压
    err = gExec.Run(exec.Command("tar", "-x", archiveFile, "-C", location))
    if err != nil {
        return err
    }

    return nil
}

// CopyFiles 复制文件与文件夹
func CopyFiles(source string, target string) (err error) {
    return gExec.Run(exec.Command("cp", "-R", "-f", source, target))
}

// ChmodList 给列表中的路径赋权
func ChmodList(pathList []string, mode os.FileMode) (err error) {
    for _, p := range pathList {
        err = os.Chmod(p, mode)
        if err != nil {
            return err
        }
    }
    return nil
}
