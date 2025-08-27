package application

import (
	"github.com/tsydim/otus-highload-architect-hw/internal/auth"
	"github.com/tsydim/otus-highload-architect-hw/internal/config"
)

type authDomain struct {
	auth *auth.AuthService
}

func buildAuthDomain(cfg *config.Config, ud usersDomain) authDomain {
	authService := auth.NewAuthService(ud.users, ud.passwords, cfg.Security)
	return authDomain{
		auth: authService,
	}
}
