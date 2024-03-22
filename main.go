package main

import (
	"fmt"
	"github.com/thewisepigeon/goo/cmd"
	"os"
)

var Version = "0.0.1"

func main() {
	var command = os.Args[1]
	if command == "version" {
		fmt.Println("Goo", Version)
		return
	}
	cmd.Execute()
}
