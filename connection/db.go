package connection

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

var db *sql.DB

func init() {
	if _, err := os.Stat("url.db"); os.IsNotExist(err) {
		os.Create("url.db")
	}
	connection, err := sql.Open("sqlite3", "file:url.db")
	if err != nil {
		panic(err)
	}
	for i, migration := range migrations {
		if _, err := connection.Exec(migration); err != nil {
			log.Println("ERROR: migration", i, err.Error())
			continue
		}
		log.Println("Migration", i, "successful")
	}
	db = connection
}

func GetDB() *sql.DB {
	return db
}
