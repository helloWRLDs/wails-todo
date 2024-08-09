package sqlite

import (
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

func Open() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite", "./todos.db")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}

func Init(db *sqlx.DB) error {
	const query = `CREATE TABLE IF NOT EXISTS "todo" (
		"id" INTEGER NOT NULL UNIQUE,
		"body" TEXT NOT NULL,
		"is_done" BOOLEAN NOT NULL DEFAULT false,
		"priority" INTEGER NOT NULL DEFAULT 5,
		"created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY("id")	
	);`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
