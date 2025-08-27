package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/tsydim/otus-highload-architect-hw/internal/apperrs"
	"github.com/tsydim/otus-highload-architect-hw/internal/auth"
	"net/http"
)

func (h *handlers) getUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	authUserId, err := auth.UserIDFromContext(ctx)
	if err != nil {
		h.handleError(ctx, w, err)
		return
	}
	userId := chi.URLParam(r, "id")
	if authUserId != userId {
		h.handleError(ctx, w, apperrs.ErrUnauthorize)
		return
	}

	user, err := h.users.Get(ctx, userId)
	if err != nil {
		h.handleError(ctx, w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}
