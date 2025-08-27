package handlers

import (
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tsydim/otus-highload-architect-hw/internal/users"

	"github.com/tsydim/otus-highload-architect-hw/internal/auth"
	"github.com/tsydim/otus-highload-architect-hw/internal/logger"
)

type handlers struct {
	auth   *auth.AuthService
	users  *users.UserService
	logger logger.Logger
}

func NewHandlers(
	auth *auth.AuthService,
	users *users.UserService,
	logger logger.Logger,
) *chi.Mux {
	r := chi.NewMux()

	h := handlers{
		auth:   auth,
		users:  users,
		logger: logger,
	}
	h.build(r)

	return r
}

func (h *handlers) build(r chi.Router) {
	r.Use(middleware.Recoverer)
	r.Post("/api/v1/user/register", h.signup)
	r.Post("/api/v1/login", h.signin)
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(h.authMiddleware)
		r.Route("/user", func(r chi.Router) {
			r.Get("/get/{id}", h.getUser)
		})
	})
}
