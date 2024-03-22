package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"os/exec"
	"strings"
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
	if os.Args[1] == "setup_db" {
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
	}
	commitMessage := os.Args[1]
	if !strings.Contains(commitMessage, "release") {
		os.Exit(0)
	}
	releaseVersion := strings.Split(commitMessage, " ")[1]
	fmt.Println(releaseVersion)
	linkerFlags := fmt.Sprintf("-X 'main.Version=%s'", releaseVersion)
	buildCmd := exec.Command("sh", "-c", fmt.Sprintf("GOOS=linux GOARCH=amd64 go build -ldflags='%s' -o goo .", linkerFlags))
	err := buildCmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
