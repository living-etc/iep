package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

type Migration struct {
	Id       string
	Filepath string
}

func (m *Migration) Run(
	ctx context.Context,
	db *sql.DB,
	logger *log.Logger,
) error {
	statements, err := os.ReadFile(m.Filepath)
	if err != nil {
		panic(err)
	}
	exec(ctx, db, string(statements))
	exec(ctx, db, "INSERT INTO migrations (id) values (?)", m.Id)

	return nil
}

func exec(ctx context.Context, db *sql.DB, statement string, args ...any) sql.Result {
	res, err := db.ExecContext(ctx, statement, args...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute statement %s: %s", statement, err)
		os.Exit(1)
	}

	return res
}
