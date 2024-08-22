package db

import (
	"database/sql"
	"fmt"
)

type Config interface {
	DriverName() string
	DSN() (string, error)
}

func Open(cfg Config) (*sql.DB, error) {
	dsn, err := cfg.DSN()
	if err != nil {
		return nil, fmt.Errorf("DSN生成エラー: %w", err)
	}
	db, err := sql.Open(cfg.DriverName(), dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open error: %w", err)
	}
	return db, nil
}
