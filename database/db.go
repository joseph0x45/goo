package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"sync"
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

func InitDB() error {
	if err := os.MkdirAll(gooHome(), 0755); err != nil {
		return err
	}
	file, err := os.Create(gooPath())
	if err != nil {
		return err
	}
	file.Close()
	db, err := sqlx.Connect("sqlite3", gooPath())
	if err != nil {
		return err
	}
	_, err = db.Exec(dbSchema)
	if err != nil {
		return err
	}
	return nil
}

var dbPool *sqlx.DB
var lock = &sync.Mutex{}

func DBConnection() *sqlx.DB {
	if dbPool == nil {
		lock.Lock()
		defer lock.Unlock()
		if dbPool == nil {
			pool, err := sqlx.Connect("sqlite3", gooPath())
			if err != nil {
				panic(err)
			}
			dbPool = pool
		}
	}
	return dbPool
}
