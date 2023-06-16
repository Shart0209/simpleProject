package main

import (
	"context"
	"errors"
	"simpleProject/internal/api"
	"simpleProject/pkg/sig"

	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

func fatal(err error, msg string) {
	if err != nil {
		log.Fatal().Err(err).Msg(msg)
	}
}

func main() {
	cfg, err := api.NewConfig()
	fatal(err, "failed get config")

	g, ctx := errgroup.WithContext(context.Background())

	svc, err := api.New(ctx, cfg)
	fatal(err, "errors create service")

	err = svc.Start(ctx, g)
	fatal(err, "errors start service")

	if err := g.Wait(); err != nil {
		if !errors.Is(err, sig.ErrShutdownSignalReceived) {
			log.Error().Err(err).Msg("errgroup errors")
		}

		log.Info().Msg("service stopping")

		err := svc.Stop()
		fatal(err, "errors stop service")

		log.Info().Msg("service stopped")
	}

}
