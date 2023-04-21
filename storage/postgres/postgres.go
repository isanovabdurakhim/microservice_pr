package postgres

import (
	"context"
	"app/config"
	"app/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db         *pgxpool.Pool
	branch storage.BranchRepoI
}

func NewPostgres(psqlConnString string, cfg config.Config) (storage.StorageI, error) {
	// First set up the pgx connection pool
	config, err := pgxpool.ParseConfig(psqlConnString)
	if err != nil {
		return nil, err
	}

	config.AfterConnect = nil
	config.MaxConns = int32(cfg.PostgresMaxConnections)

	pool, err := pgxpool.ConnectConfig(context.Background(), config)

	return &Store{
		db: pool,
	}, err
}

func (s *Store) Branch() storage.BranchRepoI {
	if s.branch == nil {
		s.branch = NewBranchRepo(s.db)
	}

	return s.branch
}
