package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/lucasmarqs/stonks/core/database/migrations"
	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
)

var datasourceName = fmt.Sprintf("%s/.stonks.db", os.Getenv("HOME"))

var db *sqlx.DB

func migrateDB() error {
	migrations := &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			&migrate.Migration{
				Id: "01",
				Up: migrations.Up_01_create_stocks,
			},
			&migrate.Migration{
				Id: "02",
				Up: migrations.Up_02_create_entries,
			},
			&migrate.Migration{
				Id: "03",
				Up: migrations.Up_03_create_reports,
			},
		},
	}

	if db == nil {
		EstablishConnection()
	}

	n, err := migrate.Exec(db.DB, "sqlite3", migrations, migrate.Up)
	if err != nil {
		return err
	}

	if n > 0 {
		fmt.Printf("applied %d migrations!\n", n)
	}
	return nil
}

// EstablishConnection attempts to create a new connection of SQLite database.
// It panics if connection or migration fail.
func EstablishConnection() *sqlx.DB {
	if db != nil {
		return db
	}

	db = sqlx.MustConnect("sqlite3", datasourceName)
	if err := migrateDB(); err != nil {
		panic(err)
	}

	return db
}
