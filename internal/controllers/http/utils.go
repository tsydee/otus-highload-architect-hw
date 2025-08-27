package handlers

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/tsydim/otus-highload-architect-hw/internal/apperrs"
	"github.com/tsydim/otus-highload-architect-hw/internal/auth"
	"net/http"
	"strings"
)

func (h *handlers) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		authHeader := r.Header.Get(tokenHeader)
		if authHeader == "" || !strings.Contains(authHeader, "Bearer ") {
			w.Header().Add("WWW-Authenticate", "Bearer")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token := strings.Replace(authHeader, "Bearer ", "", 1)
		userID, err := h.auth.Verify(token)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		ctx = auth.WithUserID(ctx, userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (h *handlers) handleError(ctx context.Context, w http.ResponseWriter, err error) {
	h.logger.With("operation", chi.RouteContext(ctx).RoutePattern()).Error(err.Error())

	switch {
	case errors.Is(err, apperrs.ErrNotFound):
		w.WriteHeader(http.StatusNotFound)
	case errors.Is(err, apperrs.ErrConditionViolation):
		w.WriteHeader(http.StatusBadRequest)
	case errors.Is(err, apperrs.ErrAlreadyExist):
		w.WriteHeader(http.StatusConflict)
	case errors.Is(err, apperrs.ErrUnauthorize):
		w.WriteHeader(http.StatusUnauthorized)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	_, err = w.Write([]byte(err.Error()))
	if err != nil {
		h.logger.Error("writer error", err.Error())
		return
	}
}
