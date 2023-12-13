package service

type Service interface {
	Install() (err error)
	Uninstall() (err error)
	Load() (err error)
	Unload() (err error)
	Reload() (err error)
	Status() (returnCode error)
}
