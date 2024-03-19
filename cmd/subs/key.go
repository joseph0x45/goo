package subs

import (
	"log"
	"os"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"github.com/thewisepigeon/goo/internal/models"
	"github.com/thewisepigeon/goo/pkg"
)

var KeyCmd = &cobra.Command{
	Use: "key",
	Run: func(cmd *cobra.Command, args []string) {
		newKey := &models.Key{
			Key: pkg.GenerateRandomString(15),
		}
		err := newKey.Save()
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
		err = clipboard.WriteAll(newKey.Key)
		if err != nil {
			log.Println("Key created: ", newKey.Key)
			os.Exit(0)
		}
		log.Println("Key created successfuly and sent into your clipboard")
		os.Exit(0)
	},
}
