package pkg

import (
	"github.com/gek64/gek/gExec"
	"os/exec"
)

func Install(pkg string) (err error) {
	return gExec.Run(exec.Command("pkg", "install", "-y", pkg))
}

func Uninstall(pkg string) (err error) {
	err = gExec.Run(exec.Command("pkg", "remove", "-y", pkg))
	if err != nil {
		return err
	}
	return gExec.Run(exec.Command("pkg", "autoremove", "-y"))
}

func Refresh() (err error) {
	return gExec.Run(exec.Command("pkg", "update"))
}

func Update(pkg string) (err error) {
	return gExec.Run(exec.Command("pkg", "upgrade", "-y", pkg))
}
