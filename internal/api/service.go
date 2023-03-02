package api

import (
	"context"
	"os"
	"simpleProject/internal/api/externalserver"
	PGStore "simpleProject/pkg/db/store"
	"simpleProject/pkg/db/store/postgres"
	libHTTP "simpleProject/pkg/http"
	"simpleProject/pkg/sig"

	"simpleProject/pkg/util"

	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

type Service interface {
	Start(ctx context.Context, g *errgroup.Group) error
	Stop() error
}

type store struct {
	pgStore PGStore.Store
	repo    PGStore.Repository
}

type service struct {
	ctx            context.Context
	cfg            *Config
	baseLogger     zerolog.Logger
	logger         zerolog.Logger
	externalServer externalserver.Server
	store          store
}

func New(ctx context.Context, cfg *Config) (Service, error) {
	logger := newLogger(cfg.LogLevel)

	svc := service{
		ctx:            ctx,
		cfg:            cfg,
		logger:         logger.With().Str("component", "api service").Logger(),
		externalServer: externalserver.New(ctx, logger.With().Str("component", "external http server").Logger()),
	}

	svc.externalServer.SetService(&svc)

	return &svc, nil
}

func (s *service) Start(ctx context.Context, g *errgroup.Group) error {
	g.Go(func() error {
		return sig.Listen(ctx)
	})

	if err := util.CreateFolder(&s.cfg.FilesFolder); err != nil {
		s.baseLogger.Error().Err(err).Send()
		return err
	}

	s.externalServer.Init(s.cfg.Listen)
	g.Go(libHTTP.MakeServerRunner(ctx,
		s.baseLogger.With().Str("component", "external_http_runner").Logger(),
		s.externalServer.GetServer()))

	pgStorage, err := postgres.NewStore(
		context.Background(),
		&postgres.Config{
			PostgresHost:           s.cfg.Postgres.Host,
			PostgresUsername:       s.cfg.Postgres.User,
			PostgresPassword:       s.cfg.Postgres.Pswd,
			PostgresDBName:         s.cfg.Postgres.DbName,
			PostgresPort:           s.cfg.Postgres.Port,
			PostgresConnectTimeout: s.cfg.Postgres.ConnectTimeout,
			PostgresMaxConns:       s.cfg.Postgres.MaxConns,
		},
		s.baseLogger.With().Str("component", "postgres db").Logger())
	if err != nil {
		return err
	}

	s.store.pgStore = pgStorage
	ex, err := pgStorage.GetExecutor()
	if err != nil {
		return err
	}
	s.store.repo = pgStorage.GetRepository(ex)

	return nil
}

func (s *service) Stop() error {
	// release resources
	if err := s.store.pgStore.Stop(); err != nil {
		return err
	}

	return nil
}

func newLogger(logLevel string) zerolog.Logger {
	level, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		level = zerolog.DebugLevel
	}

	return zerolog.New(os.Stdout).Level(level).With().Timestamp().Logger()
}
