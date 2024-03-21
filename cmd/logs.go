package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thewisepigeon/goo/cmd/subs"
)

var logCmd = &cobra.Command{
	Use: "log",
}

func init() {
	logCmd.AddCommand(subs.ListLogsCMD)
	rootCmd.AddCommand(logCmd)
}
