package postgres

import (
	"app/config"
	"app/storage"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db       *pgxpool.Pool
	branch   storage.BranchRepoI
	curier   storage.CurierRepoI
	employee storage.EmployeeRepoI
	store    storage.StoreRepoI
}

func NewPostgres(psqlConnString string, cfg config.Config) (storage.StorageI, error) {
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

func (s *Store) Curier() storage.CurierRepoI {
	if s.curier == nil {
		s.curier = NewCurierRepo(s.db)
	}

	return s.curier
}

func (s *Store) Employee() storage.EmployeeRepoI {
	if s.employee == nil {
		s.employee = NewEmployeeRepo(s.db)
	}

	return s.employee
}

func (s *Store) Store() storage.StoreRepoI {
	if s.store == nil {
		s.store = NewStoreRepo(s.db)
	}

	return s.store
}
