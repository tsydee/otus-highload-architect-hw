package databases

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/stdlib"

	"github.com/jackc/pgx"

	"github.com/tsydim/otus-highload-architect-hw/internal/config"
)

type (
	CloseFn func() error
	DB      = sql.DB
)

func NewDB(cfg *config.DB) (*sql.DB, CloseFn, error) {
	connCfg, err := pgx.ParseURI(cfg.URI)
	if err != nil {
		return nil, nil, fmt.Errorf("parse URI: %w", err)
	}
	if connCfg.Database == "" {
		return nil, nil, fmt.Errorf("database name missing in URI: %s", cfg.URI)
	}
	db := stdlib.OpenDB(connCfg)
	err = db.Ping()
	if err != nil {
		return nil, nil, fmt.Errorf("check connection: %w", err)
	}

	return db, db.Close, nil
}
