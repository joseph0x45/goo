package models

import (
	"fmt"
	"github.com/thewisepigeon/goo/database"
)

type Action struct {
	ID             int    `db:"id" json:"id"`
	Name           string `db:"name" json:"name"`
	WorkDir        string `db:"workdir" json:"workdir"`
	Command        string `db:"command" json:"command"`
	RecoverCommand string `db:"recover_command" json:"recover_command"`
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

func (a *Action) Delete(id string) error {
	pool := database.DBConnection()
	_, err := pool.Exec("delete from actions where id=$1", id)
	if err != nil {
		return fmt.Errorf("Error while deleting action: %w", err)
	}
	return nil
}

func (a *Action) List() ([]Action, error) {
	pool := database.DBConnection()
	actions := []Action{}
	err := pool.Select(&actions, "select * from actions")
	if err != nil {
		return actions, fmt.Errorf("Error while getting actions: %w", err)
	}
	return actions, nil
}

func (a *Action) GetByName(name string) (*Action, error) {
	pool := database.DBConnection()
	action := new(Action)
	err := pool.Get(action, "select * from actions where name=$1", name)
	return action, err
}

func IsValidName(name string) (bool, string) {
	if name == "" {
		return false, "Name can not be empty"
	}
	ok, err := new(Action).IsNotDuplicateName(name)
	if err != nil {
		panic(err)
	}
	if !ok {
		return false, "An action with this name already exists"
	}
	return true, ""
}
