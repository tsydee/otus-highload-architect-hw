package http

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/tsydim/otus-highload-architect-hw/internal/config"
)

type closeFn func(ctx context.Context) error

func ServeHTTP(cfg *config.HTTPConfig, handler http.Handler) (closeFn, error) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return nil, fmt.Errorf("net: %w", err)
	}

	s := http.Server{Handler: handler}

	go func() {
		err := s.Serve(l)
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	closer := func(ctx context.Context) error {
		err := s.Shutdown(ctx)
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}

	return closer, nil
}
