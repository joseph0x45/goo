package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
)

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		mux := http.NewServeMux()
		mux.Handle("GET /run/{action}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello from Goo\n"))
			return
		}))
		log.Println("Goo launched on port 8080")
		err := http.ListenAndServe(":8080", mux)
		if err != nil {
			panic(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
