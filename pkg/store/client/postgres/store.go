package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"simpleProject/internal/api"
	pgStore "simpleProject/pkg/store/client"
)

type store struct {
	db         *pgxpool.Pool
	ctx        context.Context
	logger     zerolog.Logger
	repository pgStore.Repository
}

func NewStore(ctx context.Context, cfg *api.StorageConfig, logger zerolog.Logger) (pgStore.Store, error) {
	//  Example URL
	//	postgres://jack:secret@pg.example.com:5432/mydb
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?pool_max_conns=%v&connect_timeout=%d",
		cfg.User,
		cfg.Pswd,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
		cfg.MaxConns,
		cfg.ConnectTimeout,
	)

	dbpool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("DB connection failed: %w", err)
	}

	return &store{
		db:     dbpool,
		ctx:    ctx,
		logger: logger,
	}, nil
}

func (s *store) Stop() error {
	s.db.Close()
	s.logger.Info().Msg("store stopped")
	return nil
}

func (s *store) GetExecutor() (pgStore.Executor, error) {
	return s.db, nil
}

func (s *store) GetLogger() zerolog.Logger {
	return s.logger
}

func (s *store) GetRepository(ex pgStore.Executor) pgStore.Repository {
	s.repository = NewRepository(ex, s)
	return s.repository
}
