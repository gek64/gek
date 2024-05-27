package apk

import (
	"github.com/gek64/gek/gExec"
	"os/exec"
)

func Install(pkg string) (err error) {
	return gExec.Run(exec.Command("apk", "add", pkg))
}

func Uninstall(pkg string) (err error) {
	return gExec.Run(exec.Command("apk", "del", pkg, "--purge"))
}

func Refresh() (err error) {
	return gExec.Run(exec.Command("apk", "update"))
}

func Update(pkg string) (err error) {
	return gExec.Run(exec.Command("apk", "upgrade", pkg))
}
