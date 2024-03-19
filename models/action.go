package models

import (
	"fmt"
	"github.com/thewisepigeon/goo/database"
)

type Action struct {
	ID             int    `db:"id"`
	Name           string `db:"name"`
	WorkDir        string `db:"workdir"`
	Command        string `db:"command"`
	RecoverCommand string `db:"recover_command"`
}

func (a *Action) Save() error {
	pool := database.DBConnection()
	_, err := pool.NamedExec(
		"insert into actions(name, workdir, command, recover_command) values(:name, :workdir, :command, :recover_command)",
		a,
	)
	if err != nil {
		return fmt.Errorf("Error while creating action: %w", err)
	}
	return nil
}

func (a *Action) IsNotDuplicateName(name string) (bool, error) {
	pool := database.DBConnection()
	count := 0
	err := pool.QueryRowx("select count(*) from actions where name=$1", name).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("Error while creating action: %w", err)
	}
	return count == 0, nil
}
