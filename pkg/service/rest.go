package service

import (
	"github.com/bwolf1/gin-postgres-rest-service/pkg/config"
	"github.com/bwolf1/gin-postgres-rest-service/pkg/database"
)

type Product struct {
	db *database.Database
}

func New(cfg *config.Config) (*Product, error) {
	db, err := database.New(cfg)
	if err != nil {
		return nil, err
	}

	return &Product{db}, nil
}
