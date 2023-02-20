package main

import (
	"context"
	"errors"
	"simpleProject/internal/api"
	"simpleProject/pkg/sig"

	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

func main() {
	cfg, err := api.NewConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed get config")
	}
	g, ctx := errgroup.WithContext(context.Background())

	svc, err := api.New(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("errors create service")
	}

	err = svc.Start(ctx, g)
	if err != nil {
		log.Fatal().Err(err).Msg("errors start service")
	}

	if err := g.Wait(); err != nil {
		if !errors.Is(err, sig.ErrShutdownSignalReceived) {
			log.Error().Err(err).Msg("errgroup errors")
		}

		log.Info().Msg("service stopping")

		err = svc.Stop()
		if err != nil {
			log.Fatal().Err(err).Msg("errors stop service")
		}

		log.Info().Msg("service stopped")
	}

}
