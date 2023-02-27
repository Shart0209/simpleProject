package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	pgStore "simpleProject/pkg/db/store"
	"time"
)

type store struct {
	db         *pgxpool.Pool
	ctx        context.Context
	logger     zerolog.Logger
	repository pgStore.Repository
}

type Config struct {
	PostgresHost     string `envconfig:"POSTGRES_HOST" required:"true"`
	PostgresUsername string `envconfig:"POSTGRES_USERNAME" required:"true"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	PostgresDBName   string `envconfig:"POSTGRES_DB_NAME" required:"true"`
	PostgresPort     string `envconfig:"POSTGRES_PORT" required:"true"`

	PostgresConnectTimeout time.Duration `envconfig:"POSTGRES_CONNECT_TIMEOUT" required:"true" default:"1m"`
	PostgresMaxConns       int           `envconfig:"POSTGRES_MAX_CONNS" default:"32"`
}

func NewStore(ctx context.Context, cfg *Config, logger zerolog.Logger) (pgStore.Store, error) {
	//  Example URL
	//	postgres://jack:secret@pg.example.com:5432/mydb
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?pool_max_conns=%v&connect_timeout=%d",
		cfg.PostgresUsername,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDBName,
		cfg.PostgresMaxConns,
		cfg.PostgresConnectTimeout,
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
	s.logger.Info().Msg("db stopped")
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

func (s *store) GetCtx() context.Context {
	return s.ctx
}
