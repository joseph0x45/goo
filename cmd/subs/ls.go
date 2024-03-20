package subs

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/thewisepigeon/goo/models"
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

var ListActionsCMD = &cobra.Command{
	Use: "ls",
	Run: func(cmd *cobra.Command, args []string) {
		actions, err := new(models.Action).List()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if len(actions) == 0 {
			fmt.Println("No action in database")
			os.Exit(0)
		}
		data, err := json.MarshalIndent(actions, "", "  ")
		if err != nil {
			fmt.Println("Error while displaying actions:", err.Error())
			os.Exit(1)
		}
		fmt.Println(string(data))
		os.Exit(0)
	},
}
