package vivycore

import (
	"flag"
	"fmt"
	"os"
)

// ShowArgs 显示所有的命令行传入参数
func ShowArgs() {
	for i, args := range os.Args {
		fmt.Printf("args[%d]=%s\n", i, args)
	}
}

// IsFlagPassed 判断flag是否被传入
func IsFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
