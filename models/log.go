package models

import (
	"fmt"
	"github.com/thewisepigeon/goo/database"
)

type Log struct {
	ID       int    `json:"id" db:"id"`
	Action   int    `json:"action" db:"action"`
	At       string `json:"at" db:"at"`
	Command  string `json:"command" db:"command"`
	Output   string `json:"output" db:"output"`
	ExitCode int    `json:"exit_code" db:"exit_code"`
}

func (l *Log) Save() error {
	pool := database.DBConnection()
	_, err := pool.NamedExec(
		`
    insert into logs(action, at, command, output, exit_code)
    values(:action, :at, :command, :output, :exit_code)
    `,
		l,
	)
	if err != nil {
		fmt.Println("Failed to save action execution log: ", err.Error())
		return err
	}
	return nil
}

func (l *Log) List() ([]Log, error) {
	logs := []Log{}
	pool := database.DBConnection()
	err := pool.Select(&logs, "select * from logs")
	if err != nil {
		return logs, fmt.Errorf("Error while retrieving logs: %w", err)
	}
	return logs, nil
}
