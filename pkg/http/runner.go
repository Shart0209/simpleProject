package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

func MakeServerRunner(ctx context.Context, logger zerolog.Logger, server *http.Server) func() error {
	return func() error {
		errCh := make(chan error)

		go func() {
			errCh <- server.ListenAndServe()
		}()

		logger.Info().Msg("http server has been started")

		select {
		case err := <-errCh:
			logger.Info().Msg("http server is stoping")
			return err
		case <-ctx.Done():
			shutdownCtxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*5)
			err := server.Shutdown(shutdownCtxTimeout)
			cancel()

			if errors.Is(err, nil) || errors.Is(err, http.ErrServerClosed) {
				logger.Info().Msg("http server has been stopped")
				return nil
			}
			logger.Error().Err(err).Msg("failed to stop http server")
			return fmt.Errorf("shutdowning http server: %w", err)
		}

	}
}
