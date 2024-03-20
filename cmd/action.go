package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thewisepigeon/goo/cmd/subs"
)

var actionCmd = &cobra.Command{
	Use: "action",
}

func init() {
	actionCmd.AddCommand(subs.CreateActionCMD)
	actionCmd.AddCommand(subs.RemoveActionCMD)
	actionCmd.AddCommand(subs.ListActionsCMD)
	rootCmd.AddCommand(actionCmd)
}
