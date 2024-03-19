package subs

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/thewisepigeon/goo/internal/models"
)

var ListKeysCMD = &cobra.Command{
	Use: "ls",
	Run: func(cmd *cobra.Command, args []string) {
		keys, err := new(models.Key).GetKeys()
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println("Key             ID")
		for _, key := range keys {
			fmt.Printf("%s  %d\n", key.Key, key.ID)
		}
		os.Exit(0)
	},
}
