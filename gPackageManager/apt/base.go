package apt

import (
    "github.com/gek64/gek/gExec"
    "os/exec"
)

func update() (err error) {
    return gExec.Run(exec.Command("apt", "update"))
}

func cleanLeftover() (err error) {
    return gExec.Run(exec.Command("apt", "autoremove", "-y", "--purge"))
}

func install(pkg string) (err error) {
    return gExec.Run(exec.Command("apt", "install", "-y", pkg))
}

func upgrade(pkg string) (err error) {
    return gExec.Run(exec.Command("apt", "upgrade", "-y", pkg))
}

func uninstall(pkg string) (err error) {
    err = gExec.Run(exec.Command("apt", "purge", "-y", "--autoremove", pkg))
    if err != nil {
        return err
    }
    return cleanLeftover()
}
