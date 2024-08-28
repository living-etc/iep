package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
)

type Migration interface {
	up()
	down()
}

type Migration_20240828233901 struct{}

func (migration Migration_20240828233901) up(
	ctx context.Context,
	db *sql.DB,
	args ...any,
) sql.Result {
	statement := "CREATE TABLE IF NOT EXISTS exercises(id INTEGER NOT NULL, name TEXT, description TEXT, body TEXT, PRIMARY KEY(id))"
	res, err := db.ExecContext(ctx, statement, args...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute statement %s: %s", statement, err)
		os.Exit(1)
	}

	return res
}

func (migration Migration_20240828233901) down(
	ctx context.Context,
	db *sql.DB,
	args ...any,
) sql.Result {
	statement := "DROP TABLE exercises"
	res, err := db.ExecContext(ctx, statement, args...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute statement %s: %s", statement, err)
		os.Exit(1)
	}

	return res
}
