package api

import (
	"context"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	"simpleProject/internal/api/externalserver"
)

type Service interface {
	Start(ctx context.Context, g errgroup.Group) error
	Stop()
}

type service struct {
	ctx            context.Context
	cfg            *Config
	baseLogger     zerolog.Logger
	logger         zerolog.Logger
	externalServer externalserver.Server
}

func (*service) New(ctx context.Context, cfg *Config) (*Service, error) {
	logger := newLogger(cfg.LogLever)
	svc := service{}
	return
}

func (s *service) Start(ctx context.Context, g *errgroup.Group) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) Stop() {
	//TODO implement me
	panic("implement me")
}
