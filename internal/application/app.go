package application

import (
	"context"
	"fmt"
	"time"

	"github.com/tsydim/otus-highload-architect-hw/internal/config"
	handlers "github.com/tsydim/otus-highload-architect-hw/internal/controllers/http"
	"github.com/tsydim/otus-highload-architect-hw/internal/databases"
	"github.com/tsydim/otus-highload-architect-hw/internal/logger"
	"github.com/tsydim/otus-highload-architect-hw/internal/transport/http"
)

const serverShutdownTimeout = 1 * time.Minute

func Run(ctx context.Context) error {
	cfg, err := config.Parse()
	if err != nil {
		return fmt.Errorf("parse config: %w", err)
	}

	log, err := logger.New(cfg)
	if err != nil {
		return fmt.Errorf("create logger: %w", err)
	}

	db, closeDb, err := databases.NewDB(&cfg.DB)
	if err != nil {
		return fmt.Errorf("create database: %w", err)
	}
	defer func() {
		err = closeDb()
		if err != nil {
			log.Error("close db connection: ", err.Error())
		}
	}()

	usersDomain := buildUsersDomain(db)
	authDomain := buildAuthDomain(cfg, usersDomain)

	hs := handlers.NewHandlers(authDomain.auth, usersDomain.users, log)
	stopHTTPServer, err := http.ServeHTTP(&cfg.HTTP, hs)
	if err != nil {
		return fmt.Errorf("start HTTP server: %w", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), serverShutdownTimeout)
		defer cancel()

		if err := stopHTTPServer(ctx); err != nil {
			log.Errorf("graceful shutdown: %v", err)
		}
	}()

	log.Infof("app started on port: %d", cfg.HTTP.Port)
	<-ctx.Done()

	return nil
}
