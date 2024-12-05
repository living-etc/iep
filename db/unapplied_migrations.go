package db

import (
	"context"
	"database/sql"
	"sort"

	"github.com/living-etc/iep/db/migrations"
)

const (
	getMigrationIdsSQL = "SELECT * FROM migrations"
)

func Get(
	ctx context.Context,
	db *sql.DB,
) []migrations.Migration {
	unapplied_migrations := []migrations.Migration{}

	migration_ids := []string{}
	for k := range MigrationFunctionRegistry {
		migration_ids = append(migration_ids, k)
	}
	sort.Strings(migration_ids)

	for _, migration_id := range migration_ids {

		var migration string
		if err := db.QueryRow("SELECT id FROM migrations WHERE id = ?", migration_id).Scan(&migration); err != nil {
			if err == sql.ErrNoRows {
				unapplied_migrations = append(
					unapplied_migrations,
					MigrationFunctionRegistry[migration_id](),
				)
			}
		}
	}

	return unapplied_migrations
}
