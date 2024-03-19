package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thewisepigeon/goo/cmd/subs"
)

var keyCmd = &cobra.Command{
	Use: "key",
}

func init() {
	keyCmd.AddCommand(subs.CreateKeyCMD)
	keyCmd.AddCommand(subs.ListKeysCMD)
	keyCmd.AddCommand(subs.RemoveKeyCMD)
	rootCmd.AddCommand(keyCmd)
}
