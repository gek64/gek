package apt

type Pm struct {
	Name string
}

func NewPm(name string) (p *Pm) {
	return &Pm{Name: name}
}

func (p *Pm) Install() (err error) {
	err = Refresh()
	if err != nil {
		return err
	}
	return Install(p.Name)
}

func (p *Pm) Uninstall() (err error) {
	return Uninstall(p.Name)
}

func (p *Pm) Update() (err error) {
	return Update(p.Name)
}
