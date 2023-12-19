package apt

import (
	"github.com/gek64/gek/gExec"
	"os/exec"
)

func Install(pkg string) (err error) {
	return gExec.Run(exec.Command("apt", "install", "-y", pkg))
}

func Uninstall(pkg string) (err error) {
	return gExec.Run(exec.Command("apt", "purge", "-y", "--autoremove", pkg))
}

func Refresh() (err error) {
	return gExec.Run(exec.Command("apt", "update"))
}

func Update(pkg string) (err error) {
	return gExec.Run(exec.Command("apt", "upgrade", "-y", pkg))
}
