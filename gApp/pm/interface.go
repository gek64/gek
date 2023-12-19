package pm

type Pm interface {
	Install() (err error)
	Uninstall() (err error)
	Update() (err error)
}
