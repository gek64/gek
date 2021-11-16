package main

import (
	"fmt"
	"gek_service"
)

func main() {
	s := gek_service.NewService("sshd.service", "")

	code := s.Status()
	if code != nil {
		fmt.Println(code)
	}
}
