package gek_pkgmgr_pkg

import (
	"gek_exec"
	"os/exec"
)

func update() (err error) {
	return gek_exec.Run(exec.Command("pkg", "update"))
}

func cleanLeftover() (err error) {
	return gek_exec.Run(exec.Command("pkg", "autoremove", "-y"))
}

func install(pkg string) (err error) {
	return gek_exec.Run(exec.Command("pkg", "install", "-y", pkg))
}

func upgrade(pkg string) (err error) {
	return gek_exec.Run(exec.Command("pkg", "upgrade", "-y", pkg))
}

func uninstall(pkg string) (err error) {
	err = gek_exec.Run(exec.Command("pkg", "remove", "-y", pkg))
	if err != nil {
		return err
	}
	return cleanLeftover()
}
