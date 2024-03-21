package cmd

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/thewisepigeon/goo/database"
	"github.com/thewisepigeon/goo/models"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

var pool *sqlx.DB

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		mux := http.NewServeMux()
		mux.Handle("GET /run/{action}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("Authorization")
			if apiKey == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			ok, err := new(models.Key).IsValid(apiKey)
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			action, err := new(models.Action).GetByName(r.PathValue("action"))
			if err != nil {
				if err == sql.ErrNoRows {
					w.WriteHeader(http.StatusNotFound)
					return
				}
				log.Println("Error while retrieving action:", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			commands := strings.Split(action.Command, "&&&")
			for _, command := range commands {
				cmd := exec.Command("sh", "-c", command)
				cmd.Dir = action.WorkDir
				output, err := cmd.Output()
				newLog := &models.Log{
					Action:  action.ID,
					At:      time.Now().Format("15:04 02-01-2006"),
					Command: command,
				}
				if err != nil {
					newLog.Output = err.Error()
					newLog.ExitCode = cmd.ProcessState.ExitCode()
					newLog.Save()
					log.Println("Error while running action:", err.Error())
					log.Println("Running recovery commands for action ", action.Name)
					recover_cmds := strings.Split(action.RecoverCommand, "&&&")
					for _, recover_cmd := range recover_cmds {
						cmd := exec.Command("sh", "-c", recover_cmd)
						cmd.Dir = action.WorkDir
						output, err := cmd.Output()
						newLog := &models.Log{
							Action:  action.ID,
							At:      time.UTC.String(),
							Command: command,
						}
						if err != nil {
							newLog.Output = err.Error()
							newLog.ExitCode = cmd.ProcessState.ExitCode()
							newLog.Save()
							log.Println("Error while running recovery command: ", err.Error())
							w.WriteHeader(http.StatusInternalServerError)
							return
						}
						log.Println(string(output))
					}
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				newLog.Output = string(output)
				newLog.ExitCode = cmd.ProcessState.ExitCode()
				newLog.Save()
				fmt.Println(string(output))
			}
			w.WriteHeader(http.StatusOK)
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
	pool = database.DBConnection()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
