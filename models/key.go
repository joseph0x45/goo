package models

import (
	"fmt"
	"github.com/thewisepigeon/goo/database"
)

type Key struct {
	ID  int    `db:"id"`
	Key string `db:"key"`
}

func (k *Key) Save() error {
	pool := database.DBConnection()
	_, err := pool.NamedExec(
		`insert into keys(key) values(:key)`,
		k,
	)
	if err != nil {
		return fmt.Errorf("Error while creating key: %w", err)
	}
	return nil
}

func (k *Key) GetKeys() ([]Key, error) {
	pool := database.DBConnection()
	keys := []Key{}
	err := pool.Select(&keys, "select * from keys")
	if err != nil {
		return keys, fmt.Errorf("Error while retrieving keys: %w", err)
	}
	return keys, nil
}

func (k *Key) DeleteKey(id string) error {
	pool := database.DBConnection()
	_, err := pool.Exec("delete from keys where id=$1", id)
	if err != nil {
		return fmt.Errorf("Error while deleting key: %w", err)
	}
	return nil
}
