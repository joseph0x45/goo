package models

import (
	"fmt"
	"github.com/thewisepigeon/goo/pkg"
)

type Key struct {
	ID  int    `db:"id"`
	Key string `db:"key"`
}

func (k *Key) Save() error {
	pool := pkg.DBConnection()
	_, err := pool.NamedExec(
		`insert into keys(key) values(:key)`,
		k,
	)
	if err != nil {
		return fmt.Errorf("Error while creating key: %w", err)
	}
	return nil
}
