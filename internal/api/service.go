package api

import (
	"context"
	"os"
	"simpleProject/internal/api/externalserver"
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

type service struct {
	ctx            context.Context
	cfg            *Config
	baseLogger     zerolog.Logger
	logger         zerolog.Logger
	externalServer externalserver.Server
}

func New(ctx context.Context, g *errgroup.Group, cfg *Config) (Service, error) {
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
	return nil
}

func (s *service) Stop() error {
	// release resources

	return nil
}

func newLogger(logLevel string) zerolog.Logger {
	level, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		level = zerolog.DebugLevel
	}

	return zerolog.New(os.Stdout).Level(level).With().Timestamp().Logger()
}
