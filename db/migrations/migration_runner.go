package migrations

import (
	"context"
	"database/sql"
)

type MigrationRunner struct{}

func (r *MigrationRunner) Run(
	ctx context.Context,
	db *sql.DB,
	migration Migration,
) error {
	return nil
}
