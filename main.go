package main

import (
	"fmt"
	"github.com/thewisepigeon/goo/cmd"
	"os"
)

var Version = "dev"

func main() {
	var command = os.Args[1]
	if command == "version" {
		fmt.Println("Goo version ", Version)
		return
	}
	cmd.Execute()
}
