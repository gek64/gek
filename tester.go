package main

import (
	"fmt"
	"gek_service"
)

func main() {
	s := gek_service.NewService("sshd.gek_service", "")

	code := s.Status()
	if code != nil {
		fmt.Println(code)
	}
}
