package application

import (
	"github.com/tsydim/otus-highload-architect-hw/internal/databases"
	"github.com/tsydim/otus-highload-architect-hw/internal/users"
	"github.com/tsydim/otus-highload-architect-hw/internal/users/repository"
)

type usersDomain struct {
	passwords *users.PasswordService
	users     *users.UserService
}

func buildUsersDomain(db *databases.DB) usersDomain {
	repo := repository.NewRepository(db)
	passService := users.NewPasswordService()
	userService := users.NewUserService(repo, passService)

	return usersDomain{
		passwords: passService,
		users:     userService,
	}
}
