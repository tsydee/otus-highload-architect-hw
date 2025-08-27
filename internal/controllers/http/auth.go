package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tsydim/otus-highload-architect-hw/internal/auth"
)

const tokenHeader = "Authorization"

type signUpPayload = auth.SignUpData

func (h *handlers) signup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload signUpPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		h.handleError(ctx, w, err)
		return
	}
	token, err := h.auth.SignUp(ctx, payload)
	if err != nil {
		h.handleError(ctx, w, err)
		return
	}

	w.Header().Add(tokenHeader, "Bearer "+token)
}

type signInPayload struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

func (h *handlers) signin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload signInPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		h.handleError(ctx, w, err)
		return
	}
	token, err := h.auth.SignIn(ctx, payload.Id, payload.Password)
	if err != nil {
		h.handleError(ctx, w, err)
		return
	}

	w.Header().Add(tokenHeader, "Bearer "+token)
}
