package gApp

import (
	"fmt"
	"github.com/gek64/gek/gApp/pm"
	"github.com/gek64/gek/gApp/pm/apk"
	"github.com/gek64/gek/gApp/pm/apt"
	"github.com/gek64/gek/gApp/pm/pkg"
)

type Pm struct {
	pm.Pm
}

// NewPm 新建包管理器
func NewPm(pkgManager string, pkgName string) (*Pm, error) {
	var np pm.Pm
	switch pkgManager {
	case "apt":
		np = apt.NewPm(pkgName)
	case "pkg":
		np = pkg.NewPm(pkgName)
	case "apk":
		np = apk.NewPm(pkgName)
	default:
		return nil, fmt.Errorf("no supported package manager")
	}
	return &Pm{Pm: np}, nil
}

// Install 安装包
func (p *Pm) Install() (err error) {
	return p.Pm.Install()
}

// Uninstall 卸载包
func (p *Pm) Uninstall() (err error) {
	return p.Pm.Uninstall()
}

// Update 更新包
func (p *Pm) Update() (err error) {
	return p.Pm.Update()
}
