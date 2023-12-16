package database

import (
	"database/sql"

	"github.com/bwolf1/gin-postgres-rest-service/pkg/config"
	_ "github.com/lib/pq"
)

type Database struct {
	*sql.DB
}

func New(cfg *config.Config) (*Database, error) {
	db, err := sql.Open(cfg.Database.DriverName, cfg.Database.DataSourceName)
	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
}
