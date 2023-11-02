package database

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var connection_string string = os.Getenv("MY_SQL_USER") + ":" + os.Getenv("MY_SQL_PASS") + "@tcp(db:3306)/microbloggs"

func CreateConnection() *sql.DB {
	maxTries := 3
	for i := 0; i < maxTries; i++ {
		db, err := sql.Open("mysql", connection_string)
		if err != nil {
			time.Sleep(time.Duration(time.Second * 5))
			continue
		}
		if err := db.Ping(); err == nil {
			return db
		}
		time.Sleep(time.Duration(time.Second * 5))
	}
	return nil
}
