package main

import (
	"fmt"
	"gek_github"
)

func main() {
	api, err := gek_github.NewGithubAPI("gek64/WakeOnLAN")
	if err != nil {
		return
	}

	fmt.Println(api.SearchPartsInRelease([]string{"wakeonlan", "arm64", "linux"}))
}
