package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	commitMessage := os.Args[1]
	if !strings.Contains(commitMessage, "release") {
		os.Exit(0)
	}
	releaseVersion := strings.Split(commitMessage, " ")[1]
	fmt.Println(releaseVersion)
	linkerFlags := fmt.Sprintf("-X 'main.Version=%s'", releaseVersion)
	buildCmd := exec.Command("sh", "-c", fmt.Sprintf("GOOS=linux GOARCH=amd64 go build -ldflags='%s' -o goo .", linkerFlags))
	err := buildCmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
