package apt

import "os"

type Pkg struct {
    Name     string
    Leftover []string
}

func NewPkg(name string, leftover []string) (p *Pkg) {
    return &Pkg{Name: name, Leftover: leftover}
}

func (p Pkg) Install() (err error) {
    err = update()
    if err != nil {
        return err
    }

    return install(p.Name)
}

func (p Pkg) Uninstall() (err error) {
    err = uninstall(p.Name)
    if err != nil {
        return err
    }

    for _, l := range p.Leftover {
        fileInfo, err := os.Stat(l)
        if os.IsExist(err) {
            err = os.RemoveAll(fileInfo.Name())
            if err != nil {
                return err
            }
        }
    }

    return nil
}
