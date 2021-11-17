package main

import (
	"gek_exec"
	"log"
)

func main() {
	err := gek_exec.Run("ping 127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
}
