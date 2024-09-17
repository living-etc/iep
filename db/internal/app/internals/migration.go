package internals

import (
	"context"
	"database/sql"
)

type Migration struct {
	Filepath string
}

func (m *Migration) Run(ctx context.Context, db *sql.DB) {
}
