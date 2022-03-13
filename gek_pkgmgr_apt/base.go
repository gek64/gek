package gek_pkgmgr_apt

import (
	"gek_exec"
	"os/exec"
)

func update() (err error) {
	return gek_exec.Run(exec.Command("apt", "update"))
}

func cleanLeftover() (err error) {
	return gek_exec.Run(exec.Command("apt", "autoremove", "-y", "--purge"))
}

func install(pkg string) (err error) {
	return gek_exec.Run(exec.Command("apt", "install", "-y", pkg))
}

func upgrade(pkg string) (err error) {
	return gek_exec.Run(exec.Command("apt", "upgrade", "-y", pkg))
}

func uninstall(pkg string) (err error) {
	err = gek_exec.Run(exec.Command("apt", "purge", "-y", "--autoremove", pkg))
	if err != nil {
		return err
	}
	return cleanLeftover()
}
