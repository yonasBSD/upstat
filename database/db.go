package database

import (
	"database/sql"
	"embed"
	"fmt"
	"os"

  _ "modernc.org/sqlite"
	"github.com/pressly/goose/v3"
)

var DB *sql.DB
const dbFileName = "upstat.db"


//go:embed migrations/*.sql
var embedMigrations embed.FS

func DBConnect() error {
  db_file, ok := os.LookupEnv("DB_FILE")
  if !ok {
    db_file = dbFileName
  }

	var err error
	DB, err = sql.Open("sqlite", db_file)

	if err != nil {
		return fmt.Errorf("could not connect to database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}

	if err := goose.Up(DB, "migrations"); err != nil {
		panic(err)
	}

	return nil
}
