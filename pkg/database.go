package pkg

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

var dbPool *sqlx.DB
var lock = &sync.Mutex{}

func DBConnection() *sqlx.DB {
	if dbPool == nil {
		lock.Lock()
		defer lock.Unlock()
		if dbPool == nil {
			pool, err := sqlx.Connect("sqlite3", "goo.db")
			if err != nil {
				panic(err)
			}
			dbPool = pool
		}
	}
	return dbPool
}
