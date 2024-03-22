package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

var dbSchema = `
create table if not exists keys (
  id integer primary key,
  key text not null
);

create table if not exists actions (
  id integer primary key,
  name text not null unique,
  workdir text not null,
  command text not null,
  recover_command text not null
);

create table if not exists logs (
  id integer primary key,
  action integer not null,
  at text not null,
  command text not null,
  output text not null,
  exit_code integer not null,
  FOREIGN KEY(action) REFERENCES actions(id)
);
`

func gooHome() string {
	homeDir, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/.goo", homeDir)
}

func gooPath() string {
	return fmt.Sprintf("%s/goo.db", gooHome())
}

func main() {
	switch os.Args[1] {
	case "setup_db":
		err := os.MkdirAll(gooHome(), 0755)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		file, err := os.Create(gooPath())
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		file.Close()
		db, err := sqlx.Connect("sqlite3", gooPath())
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		_, err = db.Exec(dbSchema)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println("Database initialized")
		os.Exit(0)
	case "release":
		commitMessage := os.Args[2]
		releaseVersion := strings.Split(commitMessage, " ")[1]
		payload := struct {
			TagName              string `json:"tag_name"`
			TargetCommitish      string `json:"target_commitish"`
			Name                 string `json:"name"`
			Body                 string `json:"body"`
			Draft                bool   `json:"draft"`
			Prerelease           bool   `json:"prerelease"`
			GenerateReleaseNotes bool   `json:"generate_release_notes"`
		}{
			TagName:              releaseVersion,
			TargetCommitish:      "main",
			Name:                 releaseVersion,
			Body:                 fmt.Sprintf("Goo v%s released on %s", releaseVersion, time.Now().Format("02-Jan-2006")),
			Draft:                false,
			Prerelease:           false,
			GenerateReleaseNotes: false,
		}
		data, err := json.Marshal(payload)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		req, err := http.NewRequest("POST", "https://api.github.com/repos/TheWisePigeon/goo/releases", bytes.NewBuffer(data))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		req.Header.Add("Accept", "application/vnd.github+json")
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GH_TOKEN")))
		res, err := http.DefaultClient.Do(req)
		if res.StatusCode != http.StatusCreated {
			fmt.Println(res.Status)
			os.Exit(1)
		}
		os.Exit(0)
	default:
		commitMessage := os.Args[1]
		if !strings.Contains(commitMessage, "release") {
			os.Exit(0)
		}
		releaseVersion := strings.Split(commitMessage, " ")[1]
		linkerFlags := fmt.Sprintf("-X 'main.Version=%s'", releaseVersion)
		buildCmd := exec.Command("sh", "-c", fmt.Sprintf("GOOS=linux GOARCH=amd64 go build -ldflags='%s' -o goo .", linkerFlags))
		err := buildCmd.Run()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}
	os.Exit(1)
}
