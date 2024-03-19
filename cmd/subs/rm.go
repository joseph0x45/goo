package subs

import (
	"github.com/spf13/cobra"
	"github.com/thewisepigeon/goo/models"
	"log"
	"os"
)

var RemoveKeyCMD = &cobra.Command{
	Use: "rm",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			log.Println("Missing required flag id")
			os.Exit(1)
		}
		err := new(models.Key).DeleteKey(id)
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
		log.Println("Key deleted")
		os.Exit(0)
	},
}

func init() {
	RemoveKeyCMD.Flags().String("id", "", "ID of the key to be deleted")
}
