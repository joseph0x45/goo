package main

import (
	"fmt"
	"github.com/thewisepigeon/goo/cmd"
	"os"
)

var Version = "0.0.1"

func main() {
	if len(os.Args) > 1 {
		var command = os.Args[1]
		if command == "version" {
			fmt.Println("Goo", Version)
			return
		}
	}
	cmd.Execute()
}
