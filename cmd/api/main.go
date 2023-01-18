package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"simpleProject/internal/api"
)

func main() {
	cfg, err := api.NewConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed get config")
	}
	g, ctx := errgroup.WithContext(context.Background())

	fmt.Println(ctx, cfg, g)

}
